package main

import (
	"net/http"

	"github.com/andreipimenov/zerohero/lesson78/db"
	"github.com/andreipimenov/zerohero/lesson78/server"
)

func main() {
	d := db.DB{}
	s := server.Server{
		Db: &d,
	}

	http.HandleFunc("/add", s.Add)
	http.HandleFunc("/list", s.List)
	http.ListenAndServe(":80", nil)
}
