package dictionary

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/apex/log"
)

// Dictionary represents a mapping of briefs to qwertystrokes
type Dictionary map[*Brief]string

func (d *Dictionary) add(b *Brief, definition string) {
	copy := map[*Brief]string(*d)
	copy[b] = definition
	*d = copy
}

func (d *Dictionary) MarshalJSON() ([]byte, error) {
	// Keymasks are so large that something something stack overflow?
	definitions := make(map[string]string)
	copy := map[*Brief]string(*d)
	for k, def := range copy {
		definitions[k.String()] = def
	}

	// we can't use json.Marshal because that html-escapes the > in the qwerty side
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&definitions); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadFile(filename string) (*Dictionary, error) {
	inBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var inMap map[string]string
	if err = json.Unmarshal(inBytes, &inMap); err != nil {
		return nil, err
	}
	briefMap := make(map[*Brief]string)
	for key, definition := range inMap {
		brief, err := ParseBrief(key)
		if err != nil {
			return nil, err
		}

		briefMap[brief] = definition
	}
	d := Dictionary(briefMap)
	return &d, nil
}

func (d *Dictionary) MustNotCollideWith(other *Dictionary) []error {
	errs := make([]error, 0)
	for briefA, definitionA := range map[*Brief]string(*d) {
		for briefB, definitionB := range map[*Brief]string(*other) {
			if briefA.isEqual(briefB) {
				log.WithFields(log.Fields{
					"brief":       briefA,
					"definitionA": definitionA,
					"definitionB": definitionB,
				}).Warnf("Brief collides with other dictionary")
				errs = append(errs, fmt.Errorf("Brief %s collides with other dictionary (%s vs %s)", briefA, definitionA, definitionB))
			}
		}
	}
	return errs
}
