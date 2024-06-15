package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go mysql tuitorial")

	db, err := sql.Open("mysql", "root:39063612@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//fmt.Println("Successfully Connected to Mysql database")
	insert, err := db.Query("INSERT INTO users (name, email) VALUES(?, ?)", "ELLIOT", "elliot@gmail.com")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted into user tables")

}
