package db

import (
	"encoding/json"
	"os"

	"github.com/andreipimenov/zerohero/lesson78/model"
)

type FileDB struct {
}

func (d *FileDB) Add(task model.Task) {

	tasks := d.List()

	tasks = append(tasks, task)

	b, _ := json.Marshal(tasks)

	f, _ := os.Create("filedb.json")

	f.Write(b)
	f.Close()

}

func (d *FileDB) List() []model.Task {
	b, _ := os.ReadFile("filedb.json")
	var tasks []model.Task
	json.Unmarshal(b, &tasks)
	return tasks

}
