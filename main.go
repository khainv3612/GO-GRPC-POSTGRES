package main

import (
	"Exam/proto"
	"Exam/service"
	"fmt"
)

func main() {
	fmt.Println("hello")
	//service.ConnectDB()
	log := proto.LogModel{
		ClientIp: "127.0.0.1",
		ServerIp: "127.0.0.1",
		Tags:     []string{"nhat", "oi"},
	}
	logService := service.LogManageServer{
		DB: service.ConnectDB(),
	}

	//close connect
	defer logService.DB.Close()
	logReturn, _ := service.AddLog(&logService, &log)
	fmt.Println(logReturn)
}
