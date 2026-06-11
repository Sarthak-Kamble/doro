package timer

import "time"

type Timer struct {
	Duration time.Duration
	Remaining time.Duration
}

func New(duration time.Duration) *Timer {
	return &Timer{
		Duration:  duration,
		Remaining: duration,
	}
}