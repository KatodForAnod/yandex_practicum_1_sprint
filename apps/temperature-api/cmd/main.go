package main

import (
	"os"
	"temperature-api/internal/controller"
)

func main() {
	listenAddr := ":" + os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8081"
	}

	controller.StartServer(listenAddr)
}
