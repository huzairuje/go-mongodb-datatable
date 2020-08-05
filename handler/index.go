package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
)

type dataSample struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func Index(writer http.ResponseWriter) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var data = map[string]interface{}{
		"title": "Learning Golang MongoDB",
		"name":  "Batman",
	}
	err = tmpl.Execute(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func ListData(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var sampleArray []dataSample
	jsonBytes, err := json.Marshal(sampleArray)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(jsonBytes)
}

func SampleDataJson(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	sampleData := dataSample{
		Name:    "John",
		Address: "Street USA",
	}
	jsonBytes, err := json.Marshal(sampleData)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(err.Error()))
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(jsonBytes)
}
