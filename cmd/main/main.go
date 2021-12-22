package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Task struct {
	ID        int    `json:"id"`
	TaskName  string `json:"taskName"`
	TaskDescription  string `json:"taskDescription"`
	TaskUntil string `json:"taskUntil"`
	TaskStatus bool `json:"taskStatus"`
	TaskEnd string `json:"taskEnd"`
}

func getUsers() []*Task {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM checkList")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*Task
	for results.Next() {
		var u Task
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.TaskName, &u.TaskDescription, &u.TaskUntil, &u.TaskStatus, &u.TaskEnd)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//transport.InitRoutes()
}