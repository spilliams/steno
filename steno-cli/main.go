package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	cmd := newRootCmd()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "steno",
		Short: "This is a CLI tool to help in various tasks I end up doing a lot for steno.",
	}

	cmd.AddCommand(newMergeJSONCmd())

	return cmd
}

func newMergeJSONCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "merge-json <a> <b>",
		Aliases: []string{"merge", "jsonmerge", "mergejson"},
		Args:    cobra.ExactArgs(2),
		Short:   "Merges 2 JSON files into 1.",
		Long: `Merges 2 JSON files into 1. Requires two args for the input
files. If no output file is specified, the first arg will be used as
output.

This command assumes the JSON objects will be a map of string to int. In the
event of a key collision, the values will be summed.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := getJSON(args[0])
			if err != nil {
				log.Fatal(err)
			}
			a = clean(a)
			log.Println(a)

			b, err := getJSON(args[1])
			if err != nil {
				log.Fatal(err)
			}
			b = clean(b)
			log.Println(b)

			c, err := mergeJSON(a, b)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(c)
			if outputFile == "" {
				outputFile = args[0]
			}
			return putJSON(c, outputFile)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "The output file (optional)")

	return cmd
}

func getJSON(filename string) (map[string]int, error) {
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

func clean(a map[string]int) map[string]int {
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

func mergeJSON(a, b map[string]int) (map[string]int, error) {
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

func putJSON(j map[string]int, filename string) error {
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
