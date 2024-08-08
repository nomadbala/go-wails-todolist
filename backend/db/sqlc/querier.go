// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	DeleteTask(ctx context.Context, id int64) error
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTask(ctx context.Context, id int64) (Task, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (Task, error)
}

var _ Querier = (*Queries)(nil)
