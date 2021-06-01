package main

import (
	"Exam/config"
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
	fmt.Println("Welcome!")
}
