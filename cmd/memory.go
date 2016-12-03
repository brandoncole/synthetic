package cmd

import (
    "github.com/spf13/cobra"
)

func init() {

    cmd := &cobra.Command{
        Use:   "memory",
        Short: "Simulates memory usage",
        Long: `Simulates process load`,
        Run: func(cmd *cobra.Command, args []string) {

        },
    }
    RootCmd.AddCommand(cmd)

}
