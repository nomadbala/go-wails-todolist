package domain

import db "todo-list/backend/db/sqlc"

type TaskRepository interface {
	Create(task *db.Task) error
	GetAll() ([]db.Task, error)
	GetById(id int64) (*db.Task, error)
	Update(task *db.Task) error
	Delete(id int64) error
}

type TaskService interface {
	CreateTask(task *db.Task) error
	GetAllTasks() ([]db.Task, error)
	UpdateTask(task *db.Task) error
	MarkTaskDone(id int64) error
	DeleteTask(id int64) error
}
