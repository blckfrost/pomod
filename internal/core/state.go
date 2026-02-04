package core

// timer state mainly

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
