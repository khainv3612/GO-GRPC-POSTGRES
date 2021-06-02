package main

import (
	"Exam/config"
	pb "Exam/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strings"
)

func main() {
	conn, err := grpc.Dial(config.Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed %v", err)
	}
	defer conn.Close()

	log := &pb.LogModel{
		//LogId: 76,
		//ClientIp: "127.0.0.1",
		//ServerIp: "127.0.0.1",
		Tags:     []string{"one","three"},
	}

	client := pb.NewLogManageClient(conn)

	//logReturn, _ := client.CreateLog(context.Background(), log)
	//fmt.Println(logReturn)

	logReturn, _ := client.FetchLog(context.Background(), log)
	prinResultSearch(logReturn.GetLog())
}

func prinResultSearch(lst []*pb.LogModel) {
	if len(lst) == 0 {
		fmt.Println("No result")
		return
	}
	fmt.Println("\n-------------------------------------------------------------------------------------")
	fmt.Printf("|%-10v|%-15v|%-15v|%-40v|\n", "LOG_ID", "CLIENT_IP", "SERVER_IP", "TAGS")
	fmt.Println("-------------------------------------------------------------------------------------")
	for _, log := range lst {
		fmt.Printf("|%-10v|%-15v|%-15v|%-40v|\n", log.LogId, log.ClientIp, log.ServerIp, strings.Join(log.Tags, "-"))

	}
	fmt.Println("-------------------------------------------------------------------------------------")

}
