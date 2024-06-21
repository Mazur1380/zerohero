package server

import (
	"fmt"
	"net/http"

	"github.com/andreipimenov/zerohero/lesson146/internal/store"
)

type Server struct {
	mux  *http.ServeMux
	port int
	st   *store.Store
}

func New(port int, st *store.Store) *Server {
	mux := http.NewServeMux()
	return &Server{
		mux:  mux,
		port: port,
		st:   st,
	}
}

func (s *Server) Run() {
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.mux,
	}

	s.mux.HandleFunc("/push", s.pushHandler)
	s.mux.HandleFunc("/list", s.listHandler)
	s.mux.HandleFunc("/pull", s.pullHandler)

	httpSrv.ListenAndServe()

}
