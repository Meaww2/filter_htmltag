package god

import (
	"database/sql"
	"fmt"
	"log"
)

type HTMLcontent struct {
	Site    string
	Content string
}

func Query_raw(db *sql.DB, html_ch chan HTMLcontent) {

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
	count := 1
	go func() {
		for rows.Next() {
			if count == 1000 {
				break
			}
			var html HTMLcontent

			if err := rows.Scan(&html.Site, &html.Content); err != nil {
				log.Fatalf("Query fail cause: %v", err)
				continue
			}

			html_ch <- html
			count++
			fmt.Println("Scan success:", count)
		}
	}()

}
