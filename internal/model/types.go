package model

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	State       string    `json:"state"`
	Created_at  time.Time `json:"created_at"`
	Finished_at time.Time `json:"finished_at"`
	Responsible string    `json:"responsible"`
}

type TodoList struct {
	Todos      []Todo    `json:"todos"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}
