package cmd

import (
    "github.com/spf13/cobra"
)

func init() {

    cmd := &cobra.Command{
        Use:   "network",
        Short: "Simulates network usage",
        Long: `Simulates process load`,
        Run: func(cmd *cobra.Command, args []string) {

        },
    }
    RootCmd.AddCommand(cmd)

}
