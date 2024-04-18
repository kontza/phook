/*
Copyright © 2024 Juha Ruotsalainen <juha.ruotsalainen@iki.fi>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var process int
var after string
var execute string

const PID = "process"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "phook",
	Short:   "Hook into a parent process and run commands on start and after the parent exits",
	Version: "v1.0",
	Run:     rootRunner,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&process, PID, "p", -1, "Set the process ID to hook to [default $PPID]")
	rootCmd.PersistentFlags().StringVarP(&after, "after", "a", "", "The command to run after the specified process has ended")
	rootCmd.PersistentFlags().StringVarP(&execute, "execute", "e", "", "The command to run immediately")
}
