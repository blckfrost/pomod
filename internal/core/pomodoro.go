package core

import "time"

const WorkDuration = 25 * 60

func (p *Pomodoro) Toggle() {
	now := time.Now()

	switch p.State {
	case Idle:
		p.State = Running
		p.Mode = Work
		p.Duration = WorkDuration
		p.Start = now

	case Paused:
		p.State = Running
		p.Start = now.Add(-time.Duration(p.Duration-p.PausedAt) * time.Second)
		p.PausedAt = 0

	case Running:
		remaining := p.Remaining()
		p.State = Paused
		p.PausedAt = remaining
	}
}

func (p *Pomodoro) Remaining() int {
	if p.State == Paused {
		return p.PausedAt
	}

	elapsed := int(time.Since(p.Start).Seconds())
	rem := p.Duration - elapsed
	if rem < 0 {
		return 0
	}
	return rem
}
