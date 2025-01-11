package main

import "mymodule/god"

func main() {
	// get data from database name raw_data
	data := god.Get_data()
	// Mocking
	// call filter_tag()
	god.Filter_tag(data)
	// save data to database name content
	// content must be remove tag already
	// url must be remove query param(?)
}
