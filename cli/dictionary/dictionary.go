package dictionary

import (
	"bytes"
	"encoding/json"
)

// Dictionary represents a mapping of keymasks to qwertystrokes
type Dictionary map[Keymask]string

func (d *Dictionary) add(chord Keymask, definition string) {
	copy := map[Keymask]string(*d)
	copy[chord] = definition
	*d = copy
}

func (d *Dictionary) MarshalJSON() ([]byte, error) {
	definitions := make(map[string]string)
	copy := map[Keymask]string(*d)
	for k, def := range copy {
		definitions[k.String()] = def
	}

	// we can't use json.Marshal because that html-escapes >
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&definitions); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
