package god

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func AcessDB() (*sql.DB, error) {
	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("PASSWORD_DB")
		dbname   = os.Getenv("DBNAME")
	)

	_port, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Convert string port to int fail!: %v", err)
	}
	log.Println(user, password, host, port, dbname)
	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, password, host, _port, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	fmt.Println("Successfully connected!")

	return db, err
}
