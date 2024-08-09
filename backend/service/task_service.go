package service

import (
	"fmt"
	db "todo-list/backend/db/sqlc"
	"todo-list/backend/domain"
)

type TaskService struct {
	taskRepo domain.TaskRepository
}

func NewTaskService(r domain.TaskRepository) domain.TaskService {
	return &TaskService{taskRepo: r}
}

func (s TaskService) CreateTask(task *db.Task) error {
	fmt.Println("taskservice")
	fmt.Println(task)
	return s.taskRepo.Create(task)
}

func (s TaskService) GetAllTasks() ([]db.Task, error) {
	return s.taskRepo.GetAll()
}

func (s TaskService) UpdateTask(task *db.Task) error {
	return s.taskRepo.Update(task)
}

func (s TaskService) MarkTaskDone(id int64) error {
	task, err := s.taskRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("MarkTaskDone: error getting task by id: %w", err)
	}
	if task == nil {
		return fmt.Errorf("MarkTaskDone: task not found")
	}
	task.Completed = true
	return s.taskRepo.Update(task)
}

func (s TaskService) DeleteTask(id int64) error {
	return s.taskRepo.Delete(id)
}
