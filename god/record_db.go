package god

import (
	"database/sql"
	"log"
)

type DBobj struct {
	site    string
	content string
}

func Save_record(db *sql.DB, content_ch chan DBobj, monitor_ch chan int) {

	statement := `
	INSERT INTO content (site, content)
	VALUES ($1, $2)
	`

	for {
		record := <-content_ch
		_, err := db.Exec(statement, record.site, record.content)
		if err != nil {
			log.Fatalf("Site: '%s' is error!", record.site)
			panic(err)
		} else {
			d := <-monitor_ch
			d += 1
			monitor_ch <- d
		}
	}

}
