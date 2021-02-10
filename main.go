package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_ = openDB()
}

func openDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/phantom"), nil)
	if err != nil {
		panic(err)
	}
	return db
}
