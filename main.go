package main

import (
	"log"

	"github.com/NickLand74/gRPC-server-autorization/server"
)

func main() {
	server.Run() // Просто вызываем функцию без обработки ошибки
	log.Println("Server stopped")
}
