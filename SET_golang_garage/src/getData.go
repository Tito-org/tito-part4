/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

//Include
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//getTheID = function to get the ID
func getTheID(db *sql.DB) string {
	var id string

	nb, err := db.Query("SELECT COUNT(*) FROM garage")
	checkError(err)
	defer nb.Close()
	for nb.Next() {
		err := nb.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Change the ID
	i, err := strconv.Atoi(id)
	checkError(err)
	i++
	getID := strconv.Itoa(i)
	return (getID)
}

//GetData = function to get the html query
func GetData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		methodGet(w, r)
		return
	}
	if r.Method == "POST" {
		methodPost(w, r)
	}
	if r.Method != "GET" && r.Method != "POST" {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
	if marque != "" && model != "" && fileDB != "" && checkSizefile == 0 {
		connectionDb()
	}
}
