package god

import (
	"database/sql"
	"log"
	"os"
	"time"
)

type HTMLcontent struct {
	Site    string
	Content string
	IsDone  bool
}

func Query_raw(db *sql.DB, html_ch chan HTMLcontent, worker int) {

	statement := `
	SELECT raw_data.site, raw 
	FROM raw_data 
	LEFT JOIN content 
	ON content.content IS NULL;
	`

	rows, err := db.Query(statement)

	if err != nil {
		log.Fatalf("Can't Query data cause: %v", err)
	}
	defer rows.Close()

	log.Println("<<<Start Query>>>")
	count := 1

	for rows.Next() {
		var html HTMLcontent
		html.IsDone = false
		if count == 1000 {
			html.IsDone = true
			for i := 0; i < worker; i++ {
				html_ch <- html
			}
			time.Sleep(time.Second * 60)
			db.Close()
			os.Exit(0)
			break
		}

		if err := rows.Scan(&html.Site, &html.Content); err != nil {
			log.Fatalf("Query fail cause: %v", err)
			continue
		}

		html_ch <- html
		count++
		log.Println("Scan success:", count)
	}

}
