// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package hestia

import (
	"context"
)

const createTodolist = `-- name: CreateTodolist :one
INSERT INTO todolists (
  name, created_at
) VALUES (
  $1, $2
)
RETURNING id, name, created_at
`

type CreateTodolistParams struct {
	Name      string
	CreatedAt interface{}
}

func (q *Queries) CreateTodolist(ctx context.Context, arg CreateTodolistParams) (Todolist, error) {
	row := q.db.QueryRowContext(ctx, createTodolist, arg.Name, arg.CreatedAt)
	var i Todolist
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteTodolist = `-- name: DeleteTodolist :exec
DELETE FROM todolists
WHERE id = $1
`

func (q *Queries) DeleteTodolist(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodolist, id)
	return err
}

const getTodolist = `-- name: GetTodolist :one
SELECT id, name, created_at FROM todolists
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodolist(ctx context.Context, id int64) (Todolist, error) {
	row := q.db.QueryRowContext(ctx, getTodolist, id)
	var i Todolist
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listTodolists = `-- name: ListTodolists :many
SELECT id, name, created_at FROM todolists
ORDER BY name
`

func (q *Queries) ListTodolists(ctx context.Context) ([]Todolist, error) {
	rows, err := q.db.QueryContext(ctx, listTodolists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todolist
	for rows.Next() {
		var i Todolist
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodolist = `-- name: UpdateTodolist :one
UPDATE todolists
  set name = $2,
  created_at = $3
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateTodolistParams struct {
	ID        int64
	Name      string
	CreatedAt interface{}
}

func (q *Queries) UpdateTodolist(ctx context.Context, arg UpdateTodolistParams) (Todolist, error) {
	row := q.db.QueryRowContext(ctx, updateTodolist, arg.ID, arg.Name, arg.CreatedAt)
	var i Todolist
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
