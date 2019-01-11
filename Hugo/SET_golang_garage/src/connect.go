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
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//Connect to DB / Query
func connectionDb() {
	var id string

	//Connect to the db
	fmt.Printf("Go Connect to MySQL\n")
	db, err := sql.Open(""+os.Getenv("db_type")+"", ""+os.Getenv("db_username")+":"+os.Getenv("db_password")+"@tcp("+os.Getenv("db_ip")+")/"+os.Getenv("db_name")+"")
	//db, err := sql.Open("mysql", "root:PASSWORD@tcp(172.18.12.219)/Test")
	checkError(err)

	//Call function to get the ID
	id = getTheID(db)

	//Insert the value in the DB
	rows, err := db.Query("INSERT INTO garage(id, model, marque, gazLevel, location, photo) VALUES ('" + id + "','" + model + "', '" + marque + "', '" + gazLevel + "', '" + location + "', '" + fileDB + "');")
	fmt.Printf("INSERT DONE\n")
	if err != nil {
		checkError(err)
	}
	defer rows.Close()
	defer db.Close()
}
