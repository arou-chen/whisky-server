package main

import (
	"time"
	"whisky-server/server/tcp"
)

func main() {
	server := tcp.NewServer()
	go server.Server()
	for {
		time.Sleep(1 * time.Minute)
	}
}
