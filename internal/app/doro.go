package app

import (
	"doro/internal/config"
	"doro/internal/timer"
	"time"
)

type Phase int

const (
	Work Phase = iota
	ShortBreak
	LongBreak
)

type Doro struct {
	Timer   *timer.Timer
	config  config.Config
	phase   Phase
	session int
	paused  bool
}

// func NewDoro(
// 	duration time.Duration,
// ) *Doro {

// 	return &Doro{
// 		Timer:   timer.New(duration),
// 		phase:   Work,
// 		session: 1,
// 	}

// }

func NewDoro(cfg config.Config) *Doro {
	return &Doro{
		Timer:   timer.New(cfg.WorkDuration),
		config:  cfg,
		phase:   Work,
		session: 1,
	}
}

func (p *Doro) durationForPhase() time.Duration {

	switch p.phase {

	case Work:
		return p.config.WorkDuration

	case ShortBreak:
		return p.config.ShortBreakDuration

	case LongBreak:
		return p.config.LongBreakDuration
	}

	return p.config.WorkDuration
}

// Next phase switching
// func (p *Doro) nextPhase() {

// 	switch p.phase {

// 	case Work:

// 		if p.session%4 == 0 {
// 			p.phase = LongBreak
// 		} else {
// 			p.phase = ShortBreak
// 		}

// 	case ShortBreak:
// 		p.session++
// 		p.phase = Work

// 	case LongBreak:
// 		p.session = 1
// 		p.phase = Work

// 	}

// 	p.Timer.ResetDuration(
// 		p.durationForPhase(),
// 	)

// }

func (p *Doro) nextPhase() {

	switch p.phase {

	case Work:
		p.phase = ShortBreak

	case ShortBreak:
		p.phase = Work
	}

	p.Timer.ResetDuration(
		p.durationForPhase(),
	)
}

func (p *Doro) Tick() {

	if p.paused {
		return
	}

	p.Timer.Tick()

	if p.Timer.IsFinished() {
		p.nextPhase()
	}

}

func (p *Doro) TogglePause() {

	p.paused = !p.paused

}

func (p *Doro) Reset() {

	p.Timer.Reset()

	p.paused = false

}

func (p *Doro) IsPaused() bool {

	return p.paused

}

func (p *Doro) CurrentPhase() Phase {

	return p.phase

}

func (p *Doro) Session() int {

	return p.session

}

func (p Phase) String() string {

	switch p {

	case Work:
		return "Work"

	case ShortBreak:
		return "Short Break"

	case LongBreak:
		return "Long Break"

	}

	return "Unknown"
}
