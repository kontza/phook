package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func rootRunner(cmd *cobra.Command, args []string) {
	if process < 0 {
		parent, err := os.FindProcess(os.Getppid())
		if err != nil {
			panic("Could not get the parent process!")
		}
		process = parent.Pid
	}
}
