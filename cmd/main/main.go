package main

import (
	db "Rip/pkg/database"
	. "Rip/pkg/transport"
)

func main() {
	//db.MySqlDB.Begin()
	defer db.MySqlDB.Close()
	db.InitDB()
	router := InitRoute()
	err := router.Run(":8080")
	if err != nil {
		return
	}

	//Get home page
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "It Works 3")
	//})

	//Get all task list
	//router.GET("/tasks", func(c *gin.Context) {
	//	rows, err := db.Query("SELECT id, taskName, taskDescription, taskUntil, taskStatus, taskEnd FROM checkList")
	//	defer rows.Close()
	//
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	tasks := make([]Task, 0)
	//
	//	for rows.Next() {
	//		var task Task
	//		err := rows.Scan(&task.Id, &task.TaskName, &task.TaskDescription, &task.TaskUntil, &task.TaskStatus, &task.TaskEnd)
	//		if err != nil {
	//			return
	//		}
	//		tasks = append(tasks, task)
	//	}
	//	if err = rows.Err(); err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"tasks": tasks,
	//	})
	//})

	//Create new task (Body raw JSON)
	//router.POST("/task", func(c *gin.Context) {
	//	task := Task{}
	//	err = c.Bind(&task)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	rs, err := db.Exec(
	//		"INSERT INTO checkList(taskName, taskDescription, taskUntil, taskStatus, `taskEnd`) VALUES (?, ?, ?, ?, ?)",
	//		task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	id, err := rs.LastInsertId()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	msg := fmt.Sprintf("Insert successful! new task id: %d", id)
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": msg,
	//	})
	//})

	//Update the task by id (Query Params + Body raw JSON)
	//router.PUT("/task", func(c *gin.Context) {
	//	cid := c.Request.FormValue("id")
	//	id, err := strconv.Atoi(cid)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	task := Task{Id: id}
	//	err = c.Bind(&task)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	rs, err := db.Exec(
	//		"UPDATE checkList SET taskName=?, taskDescription=?, taskUntil=?, taskStatus=?, taskEnd=? WHERE id=?",
	//		task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd, task.Id)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	ra, err := rs.RowsAffected()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	msg := fmt.Sprintf("Update task %d successful %d", task.Id, ra)
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": msg,
	//	})
	//})

	//Delete the task by id
	//router.DELETE("/task", func(c *gin.Context) {
	//	cid := c.Request.FormValue("id")
	//	id, err := strconv.Atoi(cid)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	rs, err := db.Exec("DELETE FROM checkList WHERE id=?", id)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	ra, err := rs.RowsAffected()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	msg := fmt.Sprintf("Delete task %d successful %d", id, ra)
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": msg,
	//	})
	//})
}
