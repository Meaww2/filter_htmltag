package god

import (
	"database/sql"
	"fmt"
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
		count := <-monitor_ch
		if count >= 1000 {
			count++
			monitor_ch <- count
			break
		}
		record := <-content_ch
		_, err := db.Exec(statement, record.site, record.content)
		if err != nil {
			log.Fatalf("Site: '%s' is error!", record.site)
			panic(err)
		} else {
			count++
			monitor_ch <- count
			fmt.Println("Record success:", count)
		}
	}

}
