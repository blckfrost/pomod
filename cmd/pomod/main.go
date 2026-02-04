package main

import (
	"fmt"
	"os"

	"github.com/blckfrost/pomod.git/internal/notify"
)

func main() {
	switch os.Args[1] {
	case "waybar":
		fmt.Println("okay waybar here you go")
		notify.Send("waybar up and running")
	case "ctl":
		fmt.Println("waiting for controls")
		notify.Send("controls waiting")
	case "daemon":
		fmt.Println("daemon running...")
		notify.Send("Daemon is running")

	default:
		usage()
	}
}

func usage() {
	fmt.Println("Usage: pomod daemon | waybar | ctl")
}
