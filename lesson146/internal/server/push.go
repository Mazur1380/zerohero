package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andreipimenov/zerohero/lesson146/internal/model"
	"github.com/google/uuid"
)

func (s *Server) pushHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order := &model.Order{}
	err = json.Unmarshal(b, order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order.ID = uuid.New()

	s.st.Push(order)

	fmt.Fprintf(w, "Order created with id: %s", order.ID.String())

}
