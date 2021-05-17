package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"phantom/dataLayer/daos"
	"phantom/dataLayer/dbModels"
)

func main() {
	db := openDB()
	defer db.Close()
	dao := daos.ProductImageDao{DB: db}
	productImage := dbModels.ProductImage{
		ProductId: 2,
		Url:       "some url here",
	}
	_, err := dao.AddProductImage(productImage)
	if err != nil {
		println(err.Error())
	} else {
		images, err := dao.ReadProductImages(2)
		if err != nil {
			println(err.Error())
		} else {
			for i := range images {
				println(images[i].Url)
			}
		}
	}
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/phantom")
	if err != nil {
		panic(err)
	}
	return db
}
