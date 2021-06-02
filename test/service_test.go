package main

import (
	"Exam/config"
	pb "Exam/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func Test(t *testing.T) {
	dataTest := &pb.LogModel{
		//LogId: 76,
		//ClientIp: "127.0.0.1",
		//ServerIp: "127.0.0.1",
		Tags: []string{"one", "two"},
	}

	resultWant := 1

	conn, err := grpc.Dial(config.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	defer conn.Close()
	client := pb.NewLogManageClient(conn)
	logReturn, _ := client.FetchLog(context.Background(), dataTest)
	if len(logReturn.GetLog()) != resultWant {
		t.Errorf("not pass")
	}
}
