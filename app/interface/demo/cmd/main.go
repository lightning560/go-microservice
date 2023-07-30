package main

import (
	xlog "miopkg/log"
	server "redbook/app/interface/demo/internal/http"
)

func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Panic(err.Error())
	}
}
