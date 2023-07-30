package main

import (
	xlog "miopkg/log"
	server "redbook/app/interface/feed/internal/http"
)

func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Panic(err.Error())
	}
}
