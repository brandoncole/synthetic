package cmd

import (
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "synthetic",
    Short: "Synthetic simulates operational scenarios for microservices",
    Long: `
........................................................................
:                                      __   __          __   __        :
:     **            .-----.--.--.-----|  |_|  |--.-----|  |_|__.----.  :
:   *    *          |__ --|  |  |     |   _|     |  -__|   _|  |  __|  :
:  * ---- * ---- *  |_____|___  |__|__|____|__|__|_____|____|__|____|  :
:          *    *         |_____|                                v1.0  :
:           **                                                         :
........................................................................
`,
}

func init() {
    RootCmd.Flags().Uint32P("duration", "d", 0, "Limits the duration of the simulation to a specified number of seconds")
}