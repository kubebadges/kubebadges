package main

import (
	"github.com/kubebadges/kubebadges/internal/server"
)

func main() {
	app := server.NewServer()
	if err := app.Start(); err != nil {
		panic(err)
	}
}
