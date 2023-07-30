package main

import (
	"miopkg/log"
	"redbook/app/domain/feed/internal/server"
)

// export MIO_MODE=dev
// run: go run main.go -config=config.toml
func main() {
	eng := server.NewEngine()
	if err := eng.Run(); err != nil {
		log.Error(err.Error())
	}
}
