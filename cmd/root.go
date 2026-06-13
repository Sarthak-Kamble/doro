package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showVersion bool
	version     = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "doro",
	Short: "A terminal based Pomodoro timer",

	Long: `
Doro helps you focus using the Pomodoro technique.

Examples:

  doro start
  doro start --duration 25
  doro --version
`,

	Run: func(cmd *cobra.Command, args []string) {

		if showVersion {
			fmt.Println("doro ", version)
			return
		}

		cmd.Help()
	},
}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
	}
}

func init() {

	rootCmd.Flags().
		BoolVarP(
			&showVersion,
			"version",
			"v",
			false,
			"Show application version",
		)
}