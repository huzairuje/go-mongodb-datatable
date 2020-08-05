package main

import (
	"github.com/huzairuje/go-mongodb-datatable/handler"
	"net/http"
)

func main() {
	//db := config.CreateConnectionDB()
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		handler.IndexController.Index(writer)
	})
}
