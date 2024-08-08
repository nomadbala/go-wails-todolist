-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

-- name: GetAllTasks :many
SELECT * FROM tasks
ORDER BY created_at ASC;

-- name: CreateTask :one
INSERT INTO tasks (title, completed)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateTask :one
UPDATE tasks
SET
    title = $2,
    completed = $3
WHERE id = $1
RETURNING *;