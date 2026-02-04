package notify

import "os/exec"

// logic to handle notifications

func Send(msg string) {
	exec.Command("notify-send", "Pomodoro", msg).Run()
}
