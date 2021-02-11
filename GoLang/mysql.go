package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("hi welcome to my db")
	db, err := sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/articles")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var ID int
		var Language string
		var Title string
		var Content string
		err = rows.Scan(&ID, &Language, &Title, &Content)
		if err != nil {
			panic(err)
		}
		fmt.Println(ID)
		fmt.Println(Language)
		fmt.Println(Title)
		fmt.Println(Content)
	}
	defer db.Close()
}
