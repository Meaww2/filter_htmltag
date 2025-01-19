package main

import (
	"log"
	"mymodule/god"
)

type ConfigGO struct {
}

func main() {
	// acess database
	db, err := god.AcessDB()
	if err != nil {
		log.Fatalf("Can't access database cause: %v", err)
		return
	}

	worker := 3
	buffer_in := 5
	buffer_out := 3
	html_ch := make(chan god.HTMLcontent, buffer_in)
	content_ch := make(chan god.DBobj, buffer_out)
	monitor_ch := make(chan int, worker)

	// query data
	god.Query_raw(db, html_ch)

	//
	for i := 0; i < worker; i++ {
		// call filter_tag()
		go god.Filter_tag(html_ch, content_ch)
		go god.Save_record(db, content_ch, monitor_ch)
	}

	for {

	}
	// save data to database name content
	// content must be remove tag already
	// url must be remove query param(?)
}
