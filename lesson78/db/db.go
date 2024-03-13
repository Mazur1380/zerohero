package db

import (
	"github.com/andreipimenov/zerohero/lesson78/model"
)

type DB struct {
	tasks []model.Task
}

func (d *DB) Add(task model.Task) {
	d.tasks = append(d.tasks, task)

}
func (d *DB) List() []model.Task {
	return d.tasks
}
