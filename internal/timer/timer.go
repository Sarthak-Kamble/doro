package timer

import (
	"fmt"
	"time"
)

// Struct
type Timer struct {
	Duration  time.Duration
	Remaining time.Duration
}

// Function to create a new Timer which will point to the original timer
func New(duration time.Duration) *Timer {
	return &Timer{
		Duration:  duration,
		Remaining: duration,
	}
}

// Methods in Golang

func (t *Timer) Tick() {

	if t.Remaining > 0 {
		t.Remaining -= time.Second
	}

}

func (t *Timer) IsFinished() bool {

	return t.Remaining <= 0

}

// * Formatting string method
func (t *Timer) FormatRemaining() string {

	minutes := int(t.Remaining.Minutes())

	seconds := int(t.Remaining.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
