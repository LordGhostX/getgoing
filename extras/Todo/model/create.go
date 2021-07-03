package model

import (
	"database/sql"
	"log"
)

func CreateTodo() error {
	query, err := conn.Query("INSERT INTO todo (todo) VALUES (?)", "Learn Golang")
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	if err != nil {
		return err
	}
	return nil
}
