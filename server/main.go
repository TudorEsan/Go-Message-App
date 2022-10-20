package main

import (
	"fmt"
	"net"
	"os"

	"github.com/TudorEsan/Go-Message-App/protos"
	"github.com/TudorEsan/Go-Message-App/service"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	log := hclog.Default()	

	// instance of message server
	protoMessage.RegisterMessageServiceServer(gs, service.NewMessageService())

	reflection.Register(gs)

	// create listener
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Info("Starting server on port 9092")
	gs.Serve(l)
}
