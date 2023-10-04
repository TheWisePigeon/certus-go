package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "certus",
	Short: "Runs a Certus file and displays the result",
  Long: `Usage: certus run /path/to/file.certus`,
  Run: func(cmd *cobra.Command, args []string) {
    println("Hello bozo")
  },
}
