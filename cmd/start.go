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

	RunE: func(cmd *cobra.Command, args []string) error {

		if duration <= 0 {
			return fmt.Errorf(
				"duration must be greater than zero",
			)
		}

		cfg := config.Config{
			WorkDuration: time.Duration(duration) * time.Minute,

			ShortBreakDuration: 5 * time.Minute,

			LongBreakDuration: 15 * time.Minute,
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
