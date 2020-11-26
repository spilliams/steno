package main

import (
	"fmt"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	"github.com/spilliams/steno/cli/dictionary"
	"github.com/spilliams/steno/cli/jsonfile"
)

var verbose bool

func main() {
	cobra.OnInitialize(initLogger)
	cmd := newRootCmd()
	if err := cmd.Execute(); err != nil {
		log.WithError(err).Fatal("")
	}
}

func initLogger() {
	log.SetHandler(cli.Default)
	if verbose {
		log.SetLevel(log.DebugLevel)
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

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "turn this on to get MORE")

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
				return err
			}
			a = jsonfile.Clean(a)
			log.WithField("a", a).Debug("json read")

			b, err := jsonfile.ReadFile(args[1])
			if err != nil {
				return err
			}
			b = jsonfile.Clean(b)
			log.WithField("b", b).Debug("json read")

			c, err := jsonfile.Merge(a, b)
			if err != nil {
				return err
			}
			log.WithField("c", c).Debug("json merged")
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
				return err
			}
			a = jsonfile.Clean(a)
			log.WithField("a", a).Debug("json read")
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
			stroke := "#H-F"
			log.WithField("stroke", stroke).Info("Parsing input...")
			k, err := dictionary.ParseStroke(stroke)
			if err != nil {
				return err
			}
			log.WithFields(log.Fields{
				"binary": fmt.Sprintf("%#b", k),
				"string": k.String(),
			}).Info("Parsed")

			return nil
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "dict.json", "The name to save the dictionary file as (optional)")
	return cmd
}
