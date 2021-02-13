package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Library struct {
	ID     string `json:"id"`
	Book   string `json:"book"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:lucifer007@tcp(127.0.0.1:3306)/library")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/books", getbooks).Methods("GET")
	router.HandleFunc("/books/{genre}", getbook).Methods("GET")
	http.ListenAndServe(":8000", router)
}
func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lib []Library
	result, err := db.Query("SELECT ID,Book,Author,Genre from books")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var lib2 Library
		err := result.Scan(&lib2.ID, &lib2.Book, &lib2.Author, &lib2.Genre)
		if err != nil {
			panic(err.Error())
		}
		lib = append(lib, lib2)
	}
	json.NewEncoder(w).Encode(lib)
}

func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var lib []Library
	result, err := db.Query("SELECT ID,Book,Author,Genre FROM books WHERE Genre = ?", params["genre"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var lib2 Library
	for result.Next() {
		err := result.Scan(&lib2.ID, &lib2.Book, &lib2.Author, &lib2.Genre)
		fmt.Print(lib2)
		if err != nil {
			panic(err.Error())
		}
		lib = append(lib, lib2)
	}
	json.NewEncoder(w).Encode(lib)

}
