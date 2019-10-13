package cmd

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
)

var (
	// StdWriter is the io.Writer for standard output
	StdWriter io.Writer
	// ErrWriter is the io.Writer for error output
	ErrWriter io.Writer

	rootCmd = &cobra.Command{
		Use:   "parade",
		Short: "simple SSM parameters CLI",
		Long: `Parade is a simple CLI tool for AWS SSM parameter store.
	Easy to read and writer key values in your parameter store.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("Use subcommand: keys, get, set, del")
		},
	}
)

// Execute executes the root command
func Execute(w io.Writer, e io.Writer) error {
	StdWriter = w
	ErrWriter = e
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(KeysCommand)
	rootCmd.AddCommand(GetCommand)
	rootCmd.AddCommand(SetCommand)
	rootCmd.AddCommand(DelCommand)
}
