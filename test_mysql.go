package main

import (
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:a@/test_api")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
}
