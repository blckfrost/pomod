package main

import (
	"fmt"
	"net"
	"os"

	"github.com/blckfrost/pomod.git/config"
	"github.com/blckfrost/pomod.git/internal/core"
	"github.com/blckfrost/pomod.git/internal/notify"
)

const socketPath = "/tmp/pomod.sock"

var pomo = core.New()

func runDaemon() {
	os.Remove(config.SocketPath)

	l, err := net.Listen("unix", config.SocketPath)
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Pomodoro daemon running")
	notify.Send("Daemon running")

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 64)
	n, _ := conn.Read(buf)
	cmd := string(buf[:n])

	switch cmd {
	case "toggle":
		pomo.Toggle()
	}

}

func Usage() {
	fmt.Println("Usage: pomod daemon | waybar | ctl")
}
