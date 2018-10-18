/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"os"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var enter int
var imgBackground string

func sqlQuery() int {
	var id int

	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username")+":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	checkError(err)
	query, err := db.Query("SELECT COUNT(*) FROM garage")
	checkError(err)
	defer query.Close()
	for query.Next() {
		err := query.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	checkError(err)
	if id == 0 {
		defer db.Close()
		return 84
	}
	defineAlea(id)
	defer db.Close()
	return 0
}

func selectCar() {
	fmt.Printf("%d\n", nbAlea)
	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username")+":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	checkError(err)
	clean, err := db.Query("UPDATE garage SET Book = 0")
	checkError(err)
	defer clean.Close()
	run, err := db.Query("UPDATE garage SET Book = 1 WHERE id=?", nbAlea)
	checkError(err)
	defer run.Close()
	defer db.Close()
}

func cleanDB(db *sql.DB) {
	clean, err := db.Query("UPDATE garage SET Book = 0")
	checkError(err)
	defer clean.Close()
}

func check(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username")+":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	checkError(err)
	cleanDB(db)
	enter++
	if r.URL.Path == "/go" {
		imgBackground = ConvertPicture(w, r, "../img/bg.jpg")
		imghtml := "<html>  <style> .form-img { background-image: url(data:image/jpeg;base64," + imgBackground + "); height: 195px; background-position: center; background-repeat: no-repeat;} </style> <body class=\"form-img\"> <h1 style=\"text-align: center;\"> Automatic Booking is ON</h1> 	</body> </html>"
		w.Write([]byte(imghtml))
		if enter == 1 {
			fmt.Printf("go\n")
			alea()
		}
	}
	if r.URL.Path == "/stop" {
		cleanDB(db)
		enter = 0
		os.Exit(1)
	}
}
