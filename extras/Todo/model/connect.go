package model

import (
	"database/sql"
	"fmt"
	"log"
)

var conn *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/todo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")
	conn = db
	return db
}
