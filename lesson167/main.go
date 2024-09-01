package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreipimenov/zerohero/lesson167/api"
)

func main() {

	s := NewServer()
	r := http.NewServeMux()
	h := api.HandlerFromMux(s, r)

	srv := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

type Server struct {
	Hookahs []api.Hookah
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetApiV1Hookah(w http.ResponseWriter, r *http.Request) {
	resp := s.Hookahs

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) PostApiV1Hookah(w http.ResponseWriter, r *http.Request) {
	req := api.Hookah{}
	json.NewDecoder(r.Body).Decode(&req)
	s.Hookahs = append(s.Hookahs, req)

	w.WriteHeader(http.StatusOK)

}
