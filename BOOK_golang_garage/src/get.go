/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//Get the number of line that we get in our database
func getLines(nbLine int, db *sql.DB) int {
	var id string

	nb, err := db.Query("SELECT COUNT(*) FROM garage")
	checkError(err)
	defer nb.Close()
	for nb.Next() {
		err := nb.Scan(&id)
		checkError(err)
	}
	test, err := strconv.Atoi(id)
	nbLine = test
	return (nbLine)
}

//Get the data of the database in a tab
func getTabStructFill(get *data, db *sql.DB, column string) []string {
	tab := make([]string, get.nbLine)

	for i := 1; i <= get.nbLine; i++ {
		str, err := db.Query("SELECT "+column+" FROM garage WHERE id = ?", i)
		checkError(err)
		for str.Next() {
			err := str.Scan(&tab[i-1])
			checkError(err)
		}
	}
	return (tab)
}
