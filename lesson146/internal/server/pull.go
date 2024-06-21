package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type PullRequest struct {
	Id uuid.UUID
}

func (s *Server) pullHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &PullRequest{}
	err = json.Unmarshal(b, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := s.st.Pull(req.Id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Order deleted")
}
