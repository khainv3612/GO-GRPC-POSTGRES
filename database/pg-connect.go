package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host        = "localhost"
	port        = 5432
	user        = "postgres"
	password    = "3612"
	dbname      = "test"
	search_path = "public"
)

type Loger struct {
	ID       int
	ipClient string
	ipServer string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, search_path)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	var logger Loger

	//userSql := ` SELECT log_id, client_ip, server_ip FROM "public"."LOGGING" WHERE 1=1`
	userSql := `SELECT log_id, client_ip, server_ip FROM "LOGGING" WHERE 1=1`
	err = db.QueryRow(userSql).Scan(&logger.ID, &logger.ipClient, &logger.ipServer)

	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s, welcome back!\n", logger.ipClient)
}
