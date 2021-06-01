package main

import (
	"Exam/config"
	pb "Exam/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(config.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	defer conn.Close()

	log := &pb.LogModel{
		ClientIp: "127.0.0.1",
		ServerIp: "127.0.0.1",
		Tags:     []string{"nhat", "oi"},
	}

	client := pb.NewLogManageClient(conn)

	logReturn, _ := client.CreateLog(context.Background(), log)
	fmt.Println(logReturn)

}
