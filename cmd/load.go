// Package cmd contains the command implementations for the CLI.
package cmd

import (
	"github.com/brandoncole/synthetic/resources"
	"github.com/brandoncole/synthetic/simulator"
	"github.com/spf13/cobra"
	"time"
	"os"
)

var (
	FlagCPU     *bool
	FlagDisk    *bool
	FlagMemory  *bool
	FlagNetwork *bool

	FlagCalibrationDuration  *time.Duration

	FlagLoadProfile       *string
	FlagLoadProfileMin    *int
	FlagLoadProfileMax    *int
	FlagLoadProfilePeriod *time.Duration

	FlagDuration *time.Duration
)

const (
	ProfileFlat = "flat"
	ProfileSine = "sine"
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

	FlagCalibrationDuration = cmd.Flags().Duration("cd", 10 * time.Second, "Length of time to calibrate cpu, network and disk.")

	FlagLoadProfileMin = cmd.Flags().Int("profilemin", 50, "Minimum load as a percentage of available.")
	FlagLoadProfileMax = cmd.Flags().Int("profilemax", 50, "Maximum load as a percentage of available.")
	FlagLoadProfilePeriod = cmd.Flags().Duration("profileperiod", 1 * time.Minute, "Period duration for sine profile in seconds.")

	FlagDuration = cmd.Flags().Duration("duration", 0, "Amount of time to run the load for, or infinite if 0s")

	RootCmd.AddCommand(cmd)

}

func loadCmd(cmd *cobra.Command, args []string) {

	var limiter simulator.IThroughputLimiter

	switch *FlagLoadProfile {
	case ProfileFlat:
		limiter = simulator.NewThroughputLimiterFlat(float64(*FlagLoadProfileMax) / 100.0)
	case ProfileSine:
		limiter = simulator.NewThroughputLimiterSine(float64(*FlagLoadProfileMin) / 100.0, float64(*FlagLoadProfileMax) / 100.0)
	default:
		os.Exit(1)
	}

	if *FlagCPU {
		simulator := simulator.NewThroughputSimulator(limiter, resources.ProcessorSimulation)
		simulator.Duration = *FlagDuration
		simulator.CalibrationDuration = *FlagCalibrationDuration
		simulator.PeriodDuration = *FlagLoadProfilePeriod
		simulator.Run()
	}

}
