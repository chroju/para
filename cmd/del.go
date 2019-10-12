package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/chroju/parade/ssmctl"
)

var (
	// DelCommand is the command to delete key value
	DelCommand = &cobra.Command{
		Use: "del",
		Short: "Delete key value",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := args[0]

			ssmManager, err := ssmctl.New()
			if err != nil {
				fmt.Fprintln(ErrWriter, err)
				os.Exit(1)
			}

			if err = ssmManager.DeleteParameter(key); err != nil {
				fmt.Fprintln(ErrWriter, err)
				os.Exit(1)
			}

			fmt.Fprintln(ErrWriter, "done.")
		},
	}
)

func init() {
}