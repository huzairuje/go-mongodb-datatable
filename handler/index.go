package handler

import (
	"html/template"
	"net/http"
	"path"
)

type IndexController struct {
	httpWriter http.ResponseWriter
}

func (s *IndexController) Index() {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(s.httpWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(s.httpWriter, data)
	if err != nil {
		http.Error(s.httpWriter, err.Error(), http.StatusInternalServerError)
	}
}
