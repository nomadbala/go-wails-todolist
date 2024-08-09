package main

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"
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

func parsePriority(p string) (sqlcdb.Priority, error) {
	switch p {
	case "low":
		return sqlcdb.PriorityLow, nil
	case "medium":
		return sqlcdb.PriorityMedium, nil
	case "high":
		return sqlcdb.PriorityHigh, nil
	default:
		return sqlcdb.PriorityLow, nil
	}
}

func parseDate(deadline string) (sql.NullTime, error) {
	var nullDeadline sql.NullTime
	if deadline != "" {
		parsedDeadline, err := time.Parse("2006-01-02 15:04", deadline)
		if err != nil {
			return sql.NullTime{}, fmt.Errorf("error parsing deadline: %w", err)
		}
		nullDeadline = sql.NullTime{Time: parsedDeadline, Valid: true}
	} else {
		nullDeadline = sql.NullTime{Valid: false}
	}

	return nullDeadline, nil
}

// CreateTask creates a new task
func (a *App) CreateTask(title string, completed bool, priority string, deadline string) error {
	fmt.Println(priority)
	fmt.Println(deadline)
	parsedPriority, err := parsePriority(priority)

	if err != nil {
		return fmt.Errorf("error parsing priority: %w", err)
	}

	fmt.Println(reflect.TypeOf(parsedPriority))

	parsedDeadline, err := parseDate(deadline)

	fmt.Println("----")
	fmt.Println(parsedDeadline)
	fmt.Println(reflect.TypeOf(parsedDeadline))

	fmt.Println(a.taskService.GetAllTasks())

	if err != nil {
		return fmt.Errorf("error parsing deadline: %w", err)
	}

	task := &sqlcdb.Task{
		Title:     title,
		Completed: completed,
		Priority:  parsedPriority,
		Deadline:  parsedDeadline,
	}
	return a.taskService.CreateTask(task)
}

// GetAllTasks retrieves all tasks
func (a *App) GetAllTasks() ([]sqlcdb.Task, error) {
	return a.taskService.GetAllTasks()
}

// UpdateTask updates an existing task
func (a *App) UpdateTask(id int64, title string, completed bool, priority string, deadline string) error {
	parsedPriority, err := parsePriority(priority)

	if err != nil {
		return fmt.Errorf("error parsing priority: %w", err)
	}

	parsedDeadline, err := parseDate(deadline)

	if err != nil {
		return fmt.Errorf("error parsing deadline: %w", err)
	}

	task := &sqlcdb.Task{
		ID:        id,
		Title:     title,
		Completed: completed,
		Priority:  parsedPriority,
		Deadline:  parsedDeadline,
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
