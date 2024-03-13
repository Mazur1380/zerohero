package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andreipimenov/zerohero/lesson78/model"
)

type DBInterface interface {
	Add(task model.Task)
	List() []model.Task
}

type Server struct {
	Db DBInterface
}

func (s Server) Add(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	var t model.Task
	json.Unmarshal(b, &t)
	s.Db.Add(t)
}
func (s Server) List(w http.ResponseWriter, r *http.Request) {
	tasks := s.Db.List()
	b, _ := json.Marshal(tasks)

	fmt.Fprint(w, string(b))
}
