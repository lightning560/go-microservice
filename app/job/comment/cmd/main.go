package main

import (
	"miopkg/log"
	"redbook/app/job/comment/internal/server"
)

func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		log.Panic(err.Error())
	}
}
