package cmd

import (
	"github.com/brandoncole/synthetic/resources"
	"github.com/brandoncole/synthetic/simulator"
	"github.com/spf13/cobra"
	"time"
)

var (
	FlagCPU     *bool
	FlagDisk    *bool
	FlagMemory  *bool
	FlagNetwork *bool

	FlagLoadProfile       *string
	FlagLoadProfileMin    *int
	FlagLoadProfileMax    *int
	FlagLoadProfilePeriod *time.Duration
)

func init() {

	cmd := &cobra.Command{
		Use:   "load",
		Short: "Simulates synthtic loads",
		Long:  `Simulates process load`,
		Run:   loadCmd,
		Example: `
# Runs a synthetic CPU load that utilizes 50% of the CPU
synthetic load -c -p flat --profilemax 50

# Runs a synthetic CPU load that utilizes between 0% and 50% of the CPU over 30 seconds
synthetic load -c -p sine --profilemin 0 --profilemax 50 --profileperiod 30s`,
	}

	FlagCPU = cmd.Flags().BoolP("cpu", "c", false, "Enables a synthetic CPU load")
	FlagDisk = cmd.Flags().BoolP("disk", "d", false, "Enables a synthetic disk load")
	FlagMemory = cmd.Flags().BoolP("memory", "m", false, "Enables a synthetic memory load")
	FlagNetwork = cmd.Flags().BoolP("network", "n", false, "Enables a synthetic network load")

	FlagLoadProfile = cmd.Flags().StringP("profile", "p", "flat", "Specifies the load profile [flat, sine]")

	FlagLoadProfileMin = cmd.Flags().Int("profilemin", 50, "Minimum load as a percentage of available.")
	FlagLoadProfileMax = cmd.Flags().Int("profilemax", 50, "Maximum load as a percentage of available.")
	FlagLoadProfilePeriod = cmd.Flags().Duration("profileperiod", time.Minute, "Period duration for sine profile in seconds.")

	RootCmd.AddCommand(cmd)

}

func loadCmd(cmd *cobra.Command, args []string) {

	if (*FlagCPU) {
		simulator := simulator.NewSimulator(resources.ProcessorSimulation)
		simulator.Run()
	}
}
