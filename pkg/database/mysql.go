package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MySqlDB *sql.DB

func InitDB() {
	mySqlDB, err := sql.Open("mysql", "tester:tester@tcp(db:3306)/rip_db")
	if err != nil {
		log.Fatalln(err)
	}

	MySqlDB = mySqlDB
	MySqlDB.SetMaxIdleConns(32)
	MySqlDB.SetMaxOpenConns(32)

	if err := MySqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
}
