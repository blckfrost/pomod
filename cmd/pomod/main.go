package main

import "time"

type State string

const (
	Idle    State = "IDLE"
	Running State = "RUNNING"
	Paused  State = "PAUSED"
)

type Mode string

const (
	Work  Mode = "WORK"
	Break Mode = "BREAK"
)

type Pomodoro struct {
	State    State
	Mode     Mode
	Duration int
	Start    time.Time
	PauseRem int
	Count    int
}
