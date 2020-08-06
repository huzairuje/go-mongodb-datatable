package main

import (
	"github.com/huzairuje/go-mongodb-datatable/handler"
	"net/http"
)

func main() {
	//db := config.CreateConnectionDB()
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		handler.Index(writer)
	})
	http.HandleFunc("/sample-data", func(writer http.ResponseWriter, request *http.Request) {
		handler.SampleDataJson(writer)
	})
	http.HandleFunc("/list-data", func(writer http.ResponseWriter, request *http.Request) {
		handler.ListData(writer, request)
	})
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	_ = http.ListenAndServe(":9000", nil)
}
