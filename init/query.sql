-- name: GetTodolist :one
SELECT * FROM todolists
WHERE id = $1 LIMIT 1;

-- name: ListTodolists :many
SELECT * FROM todolists
ORDER BY name;

-- name: CreateTodolist :one
INSERT INTO todolists (
  name, created_at
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTodolist :exec
DELETE FROM todolists
WHERE id = $1;

-- name: UpdateTodolist :one
UPDATE todolists
  set name = $2,
  created_at = $3
WHERE id = $1
RETURNING *;