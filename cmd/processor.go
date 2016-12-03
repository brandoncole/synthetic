package cmd

import (
	"github.com/spf13/cobra"
	"github.com/brandoncole/synthetic/resources"
	"github.com/brandoncole/synthetic/simulator"
)

func init() {

	cmd := &cobra.Command{
		Use:   "processor",
		Short: "Simulates processor usage",
		Long: `Simulates process load`,
		Run: func(cmd *cobra.Command, args []string) {

			simulator := simulator.NewSimulator(resources.ProcessorSimulation)
			simulator.Run()

		},
	}
	RootCmd.AddCommand(cmd)

}
