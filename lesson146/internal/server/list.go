package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	orders := s.st.List()

	b, err := json.Marshal(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}
