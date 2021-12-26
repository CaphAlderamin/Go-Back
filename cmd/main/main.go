package main

import (
	db "Rip/pkg/database"
	. "Rip/pkg/transport"
)

func main() {
	defer db.MySqlDB.Close()
	db.InitDB()
	router := InitRoute()
	router.Run(":8080")
}
