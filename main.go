package main

import (
	launcher "whisky-server/launcher"
	"whisky-server/server/tcp"
)

func main() {
	launcher := launcher.NewLauncher()
	launcher.AddBottle(tcp.NewServer())
	launcher.Run()
}
