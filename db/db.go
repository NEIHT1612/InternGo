package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDB() *sql.DB {
	// Connection configuration
	env := os.Getenv("APP_ENV")
	filename := ".env"
	if env == "production" {
		filename = ".env.prod"
	}
	_ = godotenv.Load(filename)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	schemename := os.Getenv("SCHEME_NAME")
	fmt.Println(schemename)
	// Build connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable",
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
