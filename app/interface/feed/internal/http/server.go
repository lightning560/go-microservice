package server

import (
	"miopkg/application"
	xlog "miopkg/log"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.NewHttpServer,
	); err != nil {
		xlog.Panic("feed httpsvr startup", xlog.Any("err", err))
	}
	return eng
}
