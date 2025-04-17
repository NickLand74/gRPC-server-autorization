package main

import (
	"NickLand74/gRPC-server-autorization.git/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
