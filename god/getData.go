package god

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "kasldlkasjf"
	dbname   = "alphadict"
)

func Get_data() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	fmt.Println("Successfully connected!")
	return ""
}

func queryMultipleRows(db *sql.DB) {
	sqlStatement := `SELECT id, name FROM your_table`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("Row scan failed: %v\n", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Rows iteration error: %v\n", err)
	}
}
