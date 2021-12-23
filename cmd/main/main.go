package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	Id              int    `json:"id" form:"id"`
	TaskName        string `json:"taskName" form:"taskName"`
	TaskDescription string `json:"taskDescription" form:"taskDescription"`
	TaskUntil       string `json:"taskUntil" form:"taskUntil"`
	TaskStatus      bool   `json:"taskStatus" form:"taskStatus"`
	TaskEnd         string `json:"taskEnd" form:"taskEnd"`
}

func main() {
	db, err := sql.Open("mysql", "tester:tester@tcp(db:3306)/rip_db")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(32)
	db.SetMaxOpenConns(32)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()

	//Get home page
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It Works 3")
	})

	//Get all task list
	router.GET("/tasks", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, taskName, taskDescription, taskUntil, taskStatus, taskEnd FROM checkList")
		defer rows.Close()

		if err != nil {
			log.Fatalln(err)
		}

		tasks := make([]Task, 0)
		for rows.Next() {
			var task Task
			err := rows.Scan(&task.Id, &task.TaskName, &task.TaskDescription, &task.TaskUntil, &task.TaskStatus, &task.TaskEnd)
			if err != nil {
				return
			}
			tasks = append(tasks, task)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
	})

	//Create new task (Body raw JSON)
	router.POST("/task", func(c *gin.Context) {
		task := Task{}
		err = c.Bind(&task)
		if err != nil {
			log.Fatalln(err)
		}

		rs, err := db.Exec(
			"INSERT INTO checkList(taskName, taskDescription, taskUntil, taskStatus, `taskEnd`) VALUES (?, ?, ?, ?, ?)",
			task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("Insert successful! new task id: %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	//Update the task by id (Query Params + Body raw JSON)
	router.PUT("/task", func(c *gin.Context) {
		cid := c.Request.FormValue("id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			log.Fatalln(err)
		}

		task := Task{Id: id}
		err = c.Bind(&task)
		if err != nil {
			log.Fatalln(err)
		}

		rs, err := db.Exec(
			"UPDATE checkList SET taskName=?, taskDescription=?, taskUntil=?, taskStatus=?, taskEnd=? WHERE id=?",
			task.TaskName, task.TaskDescription, task.TaskUntil, task.TaskStatus, task.TaskEnd, task.Id)
		if err != nil {
			log.Fatalln(err)
		}

		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("Update task %d successful %d", task.Id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	//Delete the task by id
	router.DELETE("/task", func(c *gin.Context) {
		cid := c.Request.FormValue("id")
		id, err := strconv.Atoi(cid)
		if err != nil {
			log.Fatalln(err)
		}

		rs, err := db.Exec("DELETE FROM checkList WHERE id=?", id)
		if err != nil {
			log.Fatalln(err)
		}

		ra, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}

		msg := fmt.Sprintf("Delete task %d successful %d", id, ra)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	err = router.Run(":8080")
	if err != nil {
		return
	}
}

//func getUsers() []*Task {
//	// Open up our database connection.
//	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
//
//	// if there is an error opening the connection, handle it
//	if err != nil {
//		log.Print(err.Error())
//	}
//	defer db.Close()
//
//	// Execute the query
//	results, err := db.Query("SELECT * FROM checkList")
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
//
//	var users []*Task
//	for results.Next() {
//		var u Task
//		// for each row, scan the result into our tag composite object
//		err = results.Scan(&u.ID, &u.TaskName, &u.TaskDescription, &u.TaskUntil, &u.TaskStatus, &u.TaskEnd)
//		if err != nil {
//			panic(err.Error()) // proper error handling instead of panic in your app
//		}
//
//		users = append(users, &u)
//	}
//
//	return users
//}

//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}

//func userPage(w http.ResponseWriter, r *http.Request) {
//	users := getUsers()
//
//	fmt.Println("Endpoint Hit: usersPage")
//	json.NewEncoder(w).Encode(users)
//}

//func main() {
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/users", userPage)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//
//	transport.InitRoutes()
//}
