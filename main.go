package main

import (
	"log"
	"mymodule/god"

	"github.com/joho/godotenv"
)

type ConfigGO struct {
}

func main() {
	// acess database
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Load .env fail!: %v", err)
	}

	db, err := god.AcessDB()
	if err != nil {
		log.Fatalf("Can't access database cause: %v", err)
	}

	worker := 3
	buffer_in := 5
	buffer_out := 3
	html_ch := make(chan god.HTMLcontent, buffer_in)
	content_ch := make(chan god.DBobj, buffer_out)
	// end_ch := make(chan int, 1)

	// query data
	go god.Query_raw(db, html_ch, worker)

	for i := 0; i < worker; i++ {
		// call filter_tag()
		go god.Filter_tag(html_ch, content_ch)
		go god.Save_record(db, content_ch)
	}

	// add 1000 record for 1 execute
	// stop by os.exit(0) in query_raw.go
	for {

	}

	// save data to database name content
	// content must be remove tag already
	// url must be remove query param(?)
}
