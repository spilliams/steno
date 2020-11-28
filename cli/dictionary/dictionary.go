package dictionary

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Dictionary represents a mapping of keymasks to qwertystrokes
type Dictionary map[Keymask]string

func (d *Dictionary) add(chord Keymask, definition string) {
	copy := map[Keymask]string(*d)
	copy[chord] = definition
	*d = copy
}

func (d *Dictionary) MarshalJSON() ([]byte, error) {
	// Keymasks are so large that something something stack overflow?
	definitions := make(map[string]string)
	copy := map[Keymask]string(*d)
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
	var inMap map[Keymask]string
	if err = json.Unmarshal(inBytes, &inMap); err != nil {
		return nil, err
	}
	d := Dictionary(inMap)
	return &d, nil
}
