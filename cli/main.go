package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	"github.com/spilliams/steno/cli/dictionary"
	"github.com/spilliams/steno/cli/typeyprogress"
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

	cmd.AddCommand(newMergeProgressCmd())
	cmd.AddCommand(newCleanProgressCmd())
	cmd.AddCommand(newGenerateDictionaryCmd())

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "turn this on to get MORE")

	return cmd
}

func newMergeProgressCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "merge-progress <a> <b>",
		Aliases: []string{"merge"},
		Args:    cobra.ExactArgs(2),
		Short:   "Merges 2 Typey Type progress files into 1.",
		Long: `Merges 2 Typey Type progress files into 1. Requires two args for
the input files. If no output file is specified, the first arg will be used as
output.

This command assumes the JSON objects will be a map of string to int. In the
event of a key collision, the values will be summed.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := typeyprogress.ReadFile(args[0])
			if err != nil {
				return err
			}
			a = typeyprogress.Clean(a)
			log.WithField("a", a).Debug("json read")

			b, err := typeyprogress.ReadFile(args[1])
			if err != nil {
				return err
			}
			b = typeyprogress.Clean(b)
			log.WithField("b", b).Debug("json read")

			c, err := typeyprogress.Merge(a, b)
			if err != nil {
				return err
			}
			log.WithField("c", c).Debug("json merged")
			if outputFile == "" {
				outputFile = args[0]
			}
			return typeyprogress.WriteFile(c, outputFile)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "The output file (optional)")

	return cmd
}

func newCleanProgressCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "clean-progress <a>",
		Aliases: []string{"clean"},
		Args:    cobra.ExactArgs(1),
		Short:   "Cleans a Typey Type progress file",
		Long:    "Cleans a Typey Type progress file. Trims whitespace from keys (while merging them with potential duplicates), and makes sure characters are not HTML-escaped)",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := typeyprogress.ReadFile(args[0])
			if err != nil {
				return err
			}
			a = typeyprogress.Clean(a)
			log.WithField("a", a).Debug("json read")
			return typeyprogress.WriteFile(a, outputFile)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "The output file (optional)")

	return cmd
}

func newGenerateDictionaryCmd() *cobra.Command {
	var outputFile string
	cmd := &cobra.Command{
		Use:     "generate-dictionary r.json [--output dict.json]",
		Aliases: []string{"gen-dict"},
		Args:    cobra.ExactArgs(1),
		Short:   "Generates a Plover dictionary file from a set of rules.",
		RunE: func(cmd *cobra.Command, args []string) error {
			rules, err := dictionary.ReadRulesFile(args[0])
			if err != nil {
				return err
			}
			log.WithField("rules", rules).Info("rules file read")
			if err := rules.MustBeValid(); err != nil {
				return err
			}
			log.Info("Rules are valid")

			// TODO build the rest of the dictionary

			return nil
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "dict.json", "The name to save the dictionary file as (optional)")

	return cmd
}
