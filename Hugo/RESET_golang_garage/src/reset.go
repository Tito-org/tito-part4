/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//Reset the boolean of the database (set it to 0)
func resetMysql() {
	fmt.Printf("Connection MySQL...\n")
	db, err := sql.Open(""+os.Getenv("db_type")+"", ""+os.Getenv("db_username")+":"+os.Getenv("db_password")+"@tcp("+os.Getenv("db_ip")+")/"+os.Getenv("db_name")+"")
	//db, err := sql.Open("mysql", "root:PASSWORD@tcp(172.18.12.219)/Test")
	checkError(err)
	fmt.Printf("Connection Done\n")
	query, err := db.Query("UPDATE garage SET Book = 0")
	checkError(err)
	fmt.Printf("Reset Done\n")
	defer query.Close()
	defer db.Close()
}
