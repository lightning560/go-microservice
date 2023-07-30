package main

import (
	"miopkg/log"
	"redbook/app/domain/like/internal/server"
)

// export MIO_MODE=dev
// go run main.go -config=config.toml
func main() {
	eng := server.NewEngine()
	// eng.SetGovernor("127.0.0.1:9092")
	if err := eng.Run(); err != nil {
		log.Error(err.Error())
	}
}
