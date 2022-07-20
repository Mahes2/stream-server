package main

import (
	"log"
	"os"

	"github.com/codespade/stream-server/server"
)

func main() {
	cmd := os.Args[1]
	port := os.Args[2]

	if cmd == "http" {
		if err := server.InitHttp(port); err != nil {
			log.Fatal(err)
			return
		}
	}

	if cmd == "grpc" {
		if err := server.InitGRPC(port); err != nil {
			log.Fatal(err)
			return
		}
	}

}
