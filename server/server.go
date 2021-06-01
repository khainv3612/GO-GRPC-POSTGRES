package main

import (
	"Exam/config"
	"Exam/proto"
	"Exam/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", config.PortServer)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	logService := service.LogManageServer{
		DB: service.ConnectDB(),
	}
	proto.RegisterLogManageServer(server, logService)

	fmt.Println("=====> Server started.")
	server.Serve(lis)

}
