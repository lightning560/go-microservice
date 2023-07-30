package server

import (
	"fmt"
	"miopkg/application"
	"time"

	xlog "miopkg/log"
	xcron "miopkg/worker/cron"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.startJobs,
	); err != nil {
		xlog.Panic("startup engine", xlog.Any("err", err))
	}

	return eng
}

func (eng *Engine) startJobs() error {
	cron := xcron.StdConfig("test").Build()
	cron.Schedule(xcron.Every(time.Second*10), xcron.FuncJob(eng.execJob))
	return eng.Schedule(cron)
}

func (eng *Engine) execJob() error {
	xlog.Info("info job")
	xlog.Warn("warn job")
	fmt.Println("run job")
	return nil
}
