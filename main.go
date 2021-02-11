package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_ = openDB()
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
