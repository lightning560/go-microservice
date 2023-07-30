package main

import (
	"miopkg/log"
	server "redbook/app/interface/comment/internal/http"
)

func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		log.Panic(err.Error())
	}
}
