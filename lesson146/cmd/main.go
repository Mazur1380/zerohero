package main

import (
	"github.com/andreipimenov/zerohero/lesson146/internal/server"
	"github.com/andreipimenov/zerohero/lesson146/internal/store"
)

func main() {

	st := store.New()
	port := 80

	srv := server.New(port, st)
	srv.Run()

}
