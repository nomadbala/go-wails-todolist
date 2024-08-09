-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: GetAllTasks :many
SELECT * FROM tasks
ORDER BY priority;

-- name: CreateTask :one
INSERT INTO tasks (title, completed, priority, deadline)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTask :one
UPDATE tasks
SET
    title = $2,
    completed = $3,
    priority = $4,
    deadline = $5
WHERE id = $1
RETURNING *;