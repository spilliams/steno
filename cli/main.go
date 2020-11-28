package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	cmd.AddCommand(newCompareDictionariesCmd())

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
		Long: `Generates a Plover dictionary file from a set of rules. It uses
the following factory options:
Non-standard modifier combinations: true,
Fingerspellings: true,
Left-hand numbers: Numbers 0-5
Left-hand star-numbers: F1-F5 and F12`,
		RunE: func(cmd *cobra.Command, args []string) error {
			rules, err := dictionary.ReadRulesFile(args[0])
			if err != nil {
				return err
			}
			log.WithField("rules", rules).Info("rules file read")
			if errs := rules.MustBeValid(); len(errs) > 0 {
				for _, err := range errs {
					log.Error(err.Error())
				}
				return fmt.Errorf("rules file was invalid")
			}
			log.Info("rules are valid")

			// TODO: make these factory options configurable from command line
			f := dictionary.NewFactory(dictionary.FactoryOpts{
				NonstandardModCombinations: true,
				Fingerspellings:            true,
				NumbersLeft:                dictionary.NumberOptionNumbers,
				NumberStarsLeft:            dictionary.NumberOptionFunctions,
			})
			d := f.Generate(rules)

			// encode d into a buffer of bytes
			buf := new(bytes.Buffer)
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			enc.SetIndent("", "  ")
			if err := enc.Encode(d); err != nil {
				return err
			}

			// write those bytes to a file
			log.WithField("filename", outputFile).Info("writing dictionary file")
			return ioutil.WriteFile(outputFile, buf.Bytes(), 0644)
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "dict.json", "The name to save the dictionary file as (optional)")

	return cmd
}

func newCompareDictionariesCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "compare-dictionaries a.json b.json",
		Aliases: []string{"cmp-dict"},
		Args:    cobra.ExactArgs(2),
		Short:   "Compares two dictionary files and prints out the collisions in their chords",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := dictionary.ReadFile(args[0])
			if err != nil {
				return err
			}
			b, err := dictionary.ReadFile(args[1])
			if err != nil {
				return err
			}

			errs := a.MustNotCollideWith(b)
			if len(errs) > 0 {
				log.Warn("You may want to perform lookups to see if there are other briefs for the definitions. You may also check for multi-brief combinations (such as `RE` to mean `{re^}` for a word starting in \"re\")")
			}
			return nil
		},
	}
}
