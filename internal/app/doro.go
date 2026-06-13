package app

import (
	"doro/internal/timer"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Pomodoro struct {
	Timer *timer.Timer
}

func NewDoro(duration time.Duration) *Pomodoro {
	return &Pomodoro{
		Timer: timer.New(duration),
	}
}

func (p *Pomodoro) Start() {

	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		os.Interrupt,
	)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for !p.Timer.IsFinished() {

		select {

		case <-stop:

			fmt.Println("\nSession cancelled")
			return

		case <-ticker.C:

			fmt.Println(
				p.Timer.FormatRemaining(),
			)

			p.Timer.Tick()
		}
	}

	fmt.Println("Session Completed ✨")
}
