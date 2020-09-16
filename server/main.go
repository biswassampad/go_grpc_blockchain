package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatal("Unable to listen on port 8080 : %v", err)
	}

	srv := grpc.NewServer()
}
