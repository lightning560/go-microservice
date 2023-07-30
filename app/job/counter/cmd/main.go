package main

import (
	xlog "miopkg/log"
	"redbook/app/job/counter/internal/server"
)

// go run main.go -config=config.toml
func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}
