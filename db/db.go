package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// Init initializes the database connection
func Init() {
	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal(envErr)
	}

	// Capture connection properties
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBHOST"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	// Open the database
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database!")
}
