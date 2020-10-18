package models

import "database/sql"

type TaskItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TaskItemSlice []TaskItem

func GetAllTasks(db *sql.DB) (TaskItemSlice, error) {
	row, err := db.Query("SELECT ID, Text, Completed FROM TaksItems")
	if err != nil {
		return nil, err
	}

	tasks := make(TaskItemSlice, 0, 10)
	for row.Next() {
		task := TaskItem{}
		if err := row.Scan(&task.ID, &task.Text, &task.Completed); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
