package core

import "time"

// timer ticking and the loop logic

type Pomodoro struct {
	State    State
	Mode     Mode
	Duration int
	Start    time.Time
	PauseRem int
	Count    int
}
