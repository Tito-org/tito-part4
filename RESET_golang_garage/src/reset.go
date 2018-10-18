/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"os"
)

func reset() {
	db, err := sql.Open(""+ os.Getenv("db_type") +"", ""+ os.Getenv("db_username") +":"+ os.Getenv("db_password") +"@tcp("+ os.Getenv("db_ip") +")/"+ os.Getenv("db_name") +"")
	checkError(err)
	query, err := db.Query("UPDATE garage SET Book = 0")
	checkError(err)
	fmt.Printf("RESET\n")
	defer query.Close()
	defer db.Close()
}
