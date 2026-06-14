package cmd

import (
	"doro/internal/config"
	"doro/internal/ui"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var duration int

var startCmd = &cobra.Command{

	Use:   "start",
	Short: "Start a pomodoro session",

	// Run: func(cmd *cobra.Command, args []string) {

	// 	fmt.Println("Starting Doro")

	// 	p := app.NewDoro(
	// 		time.Duration(duration) * time.Minute,
	// 	)

	// 	p.Start()

	// },

	RunE: func(cmd *cobra.Command, args []string) error {

		if duration <= 0 {
			return fmt.Errorf(
				"duration must be greater than zero",
			)
		}

		cfg := config.Config{
			WorkDuration: time.Duration(duration) * time.Minute,
		}

		return ui.Run(cfg)
	},
}

func init() {

	rootCmd.AddCommand(startCmd)

	startCmd.Flags().
		IntVarP(
			&duration,
			"duration",
			"d",
			25,
			"Duration in minutes",
		)

}
