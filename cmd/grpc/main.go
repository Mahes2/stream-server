package main

import (
	"log"
	"os"

	"github.com/codespade/stream-server/server"
)

func main() {
	port := os.Args[1]
	if err := server.InitGRPC(port); err != nil {
		log.Fatal(err)
		return
	}
}
