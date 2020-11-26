package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spilliams/steno/cli/dictionary"
	"github.com/spilliams/steno/cli/jsonfile"
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
	cmd.AddCommand(newCleanJSONCmd())
	cmd.AddCommand(newGenerateDictionaryCmd())

	return cmd
}

func newMergeJSONCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "merge-json <a> <b>",
		Aliases: []string{"merge"},
		Args:    cobra.ExactArgs(2),
		Short:   "Merges 2 JSON files into 1.",
		Long: `Merges 2 JSON files into 1. Requires two args for the input
files. If no output file is specified, the first arg will be used as
output.

This command assumes the JSON objects will be a map of string to int. In the
event of a key collision, the values will be summed.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := jsonfile.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			a = jsonfile.Clean(a)
			log.Println(a)

			b, err := jsonfile.ReadFile(args[1])
			if err != nil {
				log.Fatal(err)
			}
			b = jsonfile.Clean(b)
			log.Println(b)

			c, err := jsonfile.Merge(a, b)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(c)
			if outputFile == "" {
				outputFile = args[0]
			}
			return jsonfile.WriteFile(c, outputFile)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "The output file (optional)")

	return cmd
}

func newCleanJSONCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "clean-json <a>",
		Aliases: []string{"clean"},
		Args:    cobra.ExactArgs(1),
		Short:   "Cleans a JSON file",
		Long:    "Cleans a JSON file. Trims whitespace from keys (while merging them with potential duplicates), and makes sure characters are not HTML-escaped)",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := jsonfile.ReadFile(args[0])
			if err != nil {
				log.Fatal(err)
			}
			a = jsonfile.Clean(a)
			log.Println(a)
			return jsonfile.WriteFile(a, outputFile)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "The output file (optional)")

	return cmd
}

func newGenerateDictionaryCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "generate-dictionary --rules r.json [--output dict.json]",
		Aliases: []string{"gen-dict"},
		Short:   "Generates a Plover dictionary file from a set of rules.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("parsing STPH as a Keymask...")
			k, err := dictionary.ParseKeymask("STPH")
			if err != nil {
				return err
			}
			fmt.Printf("%#b\n", k)
			fmt.Println(k)
			return nil
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "dict.json", "The name to save the dictionary file as (optional)")
	return cmd
}
