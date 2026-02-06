package core

import "time"

type State string
type Mode string

const (
	Idle    State = "IDLE"
	Running State = "RUNNING"
	Paused  State = "PAUSED"

	Work  Mode = "WORK"
	Break Mode = "BREAK"
)

type Pomodoro struct {
	State    State
	Mode     Mode
	Duration int
	Start    time.Time
	PausedAt int
}

func New() *Pomodoro {
	return &Pomodoro{
		State: Idle,
		Mode:  Work,
	}
}
