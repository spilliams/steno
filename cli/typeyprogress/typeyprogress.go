package typeyprogress

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"sort"
	"strings"
)

func ReadFile(filename string) (map[string]int, error) {
	inBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var inJSON map[string]int
	if err = json.Unmarshal(inBytes, &inJSON); err != nil {
		return nil, err
	}
	return inJSON, nil
}

func Clean(a map[string]int) map[string]int {
	// trim space from keys
	trimmed := make(map[string]int)
	for k, v := range a {
		trimmedKey := strings.TrimSpace(k)
		existing, ok := trimmed[trimmedKey]
		if ok {
			trimmed[trimmedKey] = existing + v
		} else {
			trimmed[trimmedKey] = v
		}
	}
	return trimmed
}

func Merge(a, b map[string]int) (map[string]int, error) {
	c := make(map[string]int, len(a))
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		if prior, ok := c[k]; ok {
			c[k] = prior + v
		} else {
			c[k] = v
		}
	}
	return c, nil
}

func WriteFile(j map[string]int, filename string) error {
	var keys []string
	for k := range j {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// json loves to escape some html characters
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&j); err != nil {
		return err
	}

	return ioutil.WriteFile(filename, buf.Bytes(), 0644)
}
