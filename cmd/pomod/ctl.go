package main

import (
	"net"

	"github.com/blckfrost/pomod.git/config"
)

func runCtl(args []string) {
	if len(args) == 0 {
		Usage()
		return
	}

	conn, err := net.Dial("unix ", config.SocketPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.Write([]byte(args[0]))
}
