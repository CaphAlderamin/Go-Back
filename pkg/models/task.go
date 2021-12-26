package models

import (
	db "Rip/pkg/database"
)

type Task struct {
	Id              int    `json:"id" form:"id"`
	TaskName        string `json:"taskName" form:"taskName"`
	TaskDescription string `json:"taskDescription" form:"taskDescription"`
	TaskStatus      bool   `json:"taskStatus" form:"taskStatus"`
}

//Get all tasks list
func (task *Task) GetTasks() (tasks []Task, err error) {
	tasks = make([]Task, 0)
	rows, err := db.MySqlDB.Query("SELECT id, taskName, taskDescription, taskStatus FROM checkList")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.TaskName, &task.TaskDescription, &task.TaskStatus)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//Create new task (Body raw JSON)
func (task *Task) AddTask() (id int64, err error) {
	rs, err := db.MySqlDB.Exec(
		"INSERT INTO checkList(taskName, taskDescription, taskStatus) VALUES (?, ?, ?)",
		task.TaskName, task.TaskDescription, task.TaskStatus)
	if err != nil {
		return
	}

	id, err = rs.LastInsertId()
	return
}

//Update the task status by id (Query Params + Body raw JSON)
func (task *Task) ModTaskStatus() (resTask Task, err error) {
	_, err = db.MySqlDB.Exec(
		"UPDATE checkList SET taskStatus=? WHERE id=?",
		task.TaskStatus, task.Id)
	if err != nil {
		return
	}

	resTask, err = getResTask(task.Id)
	if err != nil {
		return
	}

	return
}

//Update the task by id (Query Params + Body raw JSON)
func (task *Task) ModTask() (resTask Task, err error) {
	_, err = db.MySqlDB.Exec(
		"UPDATE checkList SET taskName=?, taskDescription=? WHERE id=?",
		task.TaskName, task.TaskDescription, task.Id)
	if err != nil {
		return
	}

	resTask, err = getResTask(task.Id)
	if err != nil {
		return
	}

	return
}

//Delete the task by id
func (task *Task) DelTask() (rv int64, err error) {
	rs, err := db.MySqlDB.Exec("DELETE FROM checkList WHERE id=?", task.Id)
	if err != nil {
		return
	}

	rv, err = rs.RowsAffected()
	return
}

func getResTask(id int) (resTask Task, err error) {
	rows, err := db.MySqlDB.Query("SELECT id, taskName, taskDescription, taskStatus FROM checkList WHERE id=?", id)
	defer rows.Close()
	if err != nil {
		return
	}

	rows.Next()
	err = rows.Scan(&resTask.Id, &resTask.TaskName, &resTask.TaskDescription, &resTask.TaskStatus)
	if err != nil {
		return
	}

	return
}
