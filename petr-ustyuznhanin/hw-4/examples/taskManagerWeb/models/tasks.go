package models

import "database/sql"

//TaskItem ...
type TaskItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

//TaskItemSlice ...
type TaskItemSlice []TaskItem

//GetAllTasks и последующие методы под капотом у ORM
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

//Insert под капотом у ORM
func (task *TaskItem) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO TaskItems (ID, Text, Completed) VALUES (?, ?, ?)",
		task.ID, task.Text, task.Completed)

	return err
}

//Update под капотом у ORM
func (task *TaskItem) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE TaskItems SET Text = ?, Completed = ?) WHERE ID = ?",
		task.Text, task.Completed, task.ID)

	return err
}

//Delete под капотом у ORM
func (task *TaskItem) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM TaskItems WHERE ID = ?",
		task.ID)

	return err
}
