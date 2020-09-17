package models

import "database/sql"

type  TaskItem struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Completed bool `json:"completed"`
}

type TaskItemSlice []TaskItem