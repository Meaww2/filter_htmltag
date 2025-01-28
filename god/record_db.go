package god

import (
	"database/sql"
	"log"
)

type DBobj struct {
	site    string
	content string
	IsDone  bool
}

func Save_record(db *sql.DB, content_ch chan DBobj) {

	statement := `
	INSERT INTO content (site, content)
	VALUES ($1, $2)
	`

	for {

		record := <-content_ch
		if record.IsDone {
			return
		}
		_, err := db.Exec(statement, record.site, record.content)
		if err != nil {
			log.Fatalf("Site: '%s' is error!", record.site)
			panic(err)
		}
		log.Printf("Record %s success.", record.site)
	}

}
