package service

import (
	dbInfo "Exam/config"
	pb "Exam/proto"
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	//"golang.org/x/net/context"
)

type LogManageServer struct {
	DB *sql.DB
}

func (l LogManageServer) CreateLog(ctx context.Context, model *pb.LogModel) (*pb.LogModel, error) {
	panic("implement me")
}

func (l LogManageServer) FetchLog(model *pb.LogModel, server pb.LogManage_FetchLogServer) error {
	panic("implement me")
}

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.Dbname, dbInfo.Search_path)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func AddLog(s *LogManageServer, log *pb.LogModel) (*pb.LogModel, error) {
	_, err := s.DB.Exec(`INSERT INTO "LOGGING"(client_ip,server_ip,tags) VALUES($1,$2,$3)`, log.ClientIp, log.ServerIp, pq.Array(log.Tags))
	if err != nil {
		return nil, err
	}
	return log, nil
}
