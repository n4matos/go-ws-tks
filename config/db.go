package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	//"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {

	/*err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}*/

	LoadEnv()

	//Getting .env values
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, dbPort, username, dbName, password) //Build connection string

	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	//defer db.Close()

	return db
}
