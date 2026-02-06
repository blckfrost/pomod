package main

import (
	"os"

	"github.com/blckfrost/pomod.git/internal/waybar"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "daemon":
		runDaemon()
	case "ctl":
		runCtl(os.Args[2:])
	case "waybar":
		waybar.Print(pomo)
	}
}
