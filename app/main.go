package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ecosia/go-prometheus-workshop/app/fetch"
)

const (
	port = "8000"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/withResponse.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	treeData, statusCode, err := fetch.Fetch(fetch.NewRequest)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	err = t.Execute(w, treeData.Count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	fmt.Printf("Service started at %v", port)
	http.ListenAndServe("0.0.0.0:"+port, mux)
}
