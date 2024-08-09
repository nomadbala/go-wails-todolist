package repository

import (
	"context"
	"database/sql"
	"fmt"
	sqlcdb "todo-list/backend/db/sqlc"
	"todo-list/backend/domain"
)

type TaskRepository struct {
	queries *sqlcdb.Queries
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	q := sqlcdb.New(db)
	return &TaskRepository{queries: q}
}

func (r TaskRepository) Create(task *sqlcdb.Task) error {
	params := sqlcdb.CreateTaskParams{
		Title:     task.Title,
		Completed: task.Completed,
		Priority:  task.Priority,
		Deadline:  task.Deadline,
	}

	_, err := r.queries.CreateTask(context.Background(), params)

	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	return nil
}

func (r TaskRepository) GetAll() ([]sqlcdb.Task, error) {
	tasks, err := r.queries.GetAllTasks(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}

	return tasks, nil
}

func (r TaskRepository) GetById(id int64) (*sqlcdb.Task, error) {
	task, err := r.queries.GetTask(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	return &task, nil
}

func (r TaskRepository) Update(task *sqlcdb.Task) error {
	params := sqlcdb.UpdateTaskParams{
		ID:        task.ID,
		Title:     task.Title,
		Completed: task.Completed,
		Priority:  task.Priority,
		Deadline:  task.Deadline,
	}

	_, err := r.queries.UpdateTask(context.Background(), params)

	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

func (r TaskRepository) Delete(id int64) error {
	if err := r.queries.DeleteTask(context.Background(), id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
