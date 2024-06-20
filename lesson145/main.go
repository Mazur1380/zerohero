package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmp *template.Template

func main() {
	var err error
	tmp, err = template.ParseFiles("index.tmpl", "about.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":80", nil)
}

type TmpData struct {
	Name     string
	Benefits []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := TmpData{
		Name: "Djas",
		Benefits: []string{
			"Premium Sport TV Channels",
			"Cashback on sport stores 5%",
			"Free entrance 1 match per year",
		},
	}
	tmp.ExecuteTemplate(w, "index", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmp.ExecuteTemplate(w, "about", nil)
}
