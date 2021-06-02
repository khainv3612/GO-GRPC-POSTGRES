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
	"strconv"
)

type LogManageServer struct {
	DB *sql.DB
}

func (s LogManageServer) CreateLog(ctx context.Context, log *pb.LogModel) (*pb.LogModel, error) {
	_, err := s.DB.Exec(`INSERT INTO "LOGGING"(client_ip,server_ip,tags) VALUES($1,$2,$3)`, log.ClientIp, log.ServerIp, pq.Array(log.Tags))
	if err != nil {
		return nil, err
	}
	return log, nil
}

func (s LogManageServer) FetchLog(ctx context.Context, model *pb.LogModel) (*pb.LogModels, error) {
	sql := `SELECT log_id,client_ip,server_ip,tags::text[] FROM "LOGGING" WHERE 1=1 `

	if &model.LogId != nil && model.LogId != 0 {
		//sql += ` AND log_id = $1 `
		sql += " AND log_id = " + strconv.Itoa(int(model.LogId))
	}
	if &model.ClientIp != nil && model.ClientIp != "" {
		//sql += ` AND client_ip = $2 `
		sql += " AND client_ip = '" + model.ClientIp + "'"
	}
	if &model.ServerIp != nil && model.ServerIp != "" {
		//sql += ` AND server_ip = $3 `
		sql += " AND server_ip = '" + model.ServerIp + "'"
	}
	if &model.Tags != nil && len(model.Tags) > 0 {
		//sql += ` AND tags @> = $4`
		var tags string
		tags = string("'{")
		index := 0
		for _, tag := range model.Tags {
			tags += tag
			index++
			if index < len(model.Tags) {
				tags += string(",")
			}
		}
		tags += string("}'")
		sql += " AND tags @> " + tags
	}
	fmt.Println(sql)
	query, err := s.DB.Query(sql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer query.Close()
	var result []*pb.LogModel
	for query.Next() {
		log := pb.LogModel{}
		err := query.Scan(&log.LogId, &log.ClientIp, &log.ServerIp, pq.Array(&log.Tags))
		if err != nil {
			panic(err)
		}
		result = append(result, &log)
	}
	logs := pb.LogModels{
		Log: result,
	}
	return &logs, nil
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
