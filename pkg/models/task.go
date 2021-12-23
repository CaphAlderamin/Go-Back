package models

import (
	db "Rip/pkg/database"
)

type Task struct {
	Id              int    `json:"id" form:"id"`
	TaskName        string `json:"taskName" form:"taskName"`
	TaskDescription string `json:"taskDescription" form:"taskDescription"`
	TaskUntil       string `json:"taskUntil" form:"taskUntil"`
	TaskStatus      bool   `json:"taskStatus" form:"taskStatus"`
	TaskEnd         string `json:"taskEnd" form:"taskEnd"`
}

//Get all tasks list
func (task *Task) GetTasks() (tasks []Task, err error) {
	tasks = make([]Task, 0)
	rows, err := db.MySqlDB.Query("SELECT id, taskName, taskDescription, taskUntil, taskStatus, taskEnd FROM checkList")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.TaskName, &task.TaskDescription, &task.TaskUntil, &task.TaskStatus, &task.TaskEnd)
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
		"INSERT INTO checkList(taskName, taskDescription, taskUntil, taskStatus, `taskEnd`) VALUES (?, ?, ?, ?, ?)",
		task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd)
	if err != nil {
		return
	}

	id, err = rs.LastInsertId()
	return
}

//Update the task by id (Query Params + Body raw JSON)
func (task *Task) ModTask() (rv int64, err error) {
	rs, err := db.MySqlDB.Exec(
		"UPDATE checkList SET taskName=?, taskDescription=?, taskUntil=?, taskStatus=?, taskEnd=? WHERE id=?",
		task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd, task.Id)
	if err != nil {
		return
	}

	rv, err = rs.RowsAffected()
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
