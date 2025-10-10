package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB() *sql.DB {
	// Connection configuration
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "123456"
	dbname := "manage_goods"
	schemename := "goods_management"

	// Build connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schemename)

	// Open connection
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("Connection failed!")
	}

	log.Println("Connection successful.")
	return DB
}
