package main

import (
	"context"
	"fmt"
	sqlcdb "todo-list/backend/db/sqlc"
	"todo-list/backend/domain"
	"todo-list/backend/service"
)

// App struct
type App struct {
	ctx         context.Context
	taskService domain.TaskService
}

// NewApp creates a new App application struct
func NewApp(taskRepo domain.TaskRepository) *App {
	taskService := service.NewTaskService(taskRepo)
	return &App{taskService: taskService}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// CreateTask creates a new task
func (a *App) CreateTask(title string, completed bool) error {
	task := &sqlcdb.Task{
		Title:     title,
		Completed: completed,
	}
	return a.taskService.CreateTask(task)
}

// GetAllTasks retrieves all tasks
func (a *App) GetAllTasks() ([]sqlcdb.Task, error) {
	return a.taskService.GetAllTasks()
}

// UpdateTask updates an existing task
func (a *App) UpdateTask(id int64, title string, completed bool) error {
	task := &sqlcdb.Task{
		ID:        id,
		Title:     title,
		Completed: completed,
	}
	return a.taskService.UpdateTask(task)
}

// MarkTaskDone marks a task as completed
func (a *App) MarkTaskDone(id int64) error {
	return a.taskService.MarkTaskDone(id)
}

// DeleteTask deletes a task
func (a *App) DeleteTask(id int64) error {
	return a.taskService.DeleteTask(id)
}
