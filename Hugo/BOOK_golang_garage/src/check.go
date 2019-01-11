/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"database/sql"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var repetition int
var str string

//Check the err in our method
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Add a space in our string
func addSpace(one string, two string) {
	str = one
	for i := 0; i <= repetition; i++ {
		str = str + " "
	}
	str = str + two
	repetition = 0
}

//Check if our string get a "%" (= space symbol in URL language)
func checkString(url string) string {
	i := strings.Index(url, "%")
	if i > -1 {
		one := url[:i]
		two := url[i+3:]
		if url[i+3] == '%' {
			repetition++
			return (checkString(one + two))
		}
		if repetition > 0 {
			addSpace(one, two)
			return (str)
		}
		str = one + " " + two
		return (str)
	}
	return (url)
}

//Check if the Element is book or not
func checkBooking(i int) int {
	var booking string

	db, err := sql.Open(""+os.Getenv("db_type")+"", ""+os.Getenv("db_username")+":"+os.Getenv("db_password")+"@tcp("+os.Getenv("db_ip")+")/"+os.Getenv("db_name")+"")
	//db, err := sql.Open("mysql", "root:PASSWORD@tcp(172.18.12.219)/Test")
	checkError(err)
	query, err := db.Query("SELECT Book FROM garage WHERE id=?", i+1)
	checkError(err)
	for query.Next() {
		err := query.Scan(&booking)
		checkError(err)
	}
	book, err := strconv.Atoi(booking)
	checkError(err)
	if book == 1 {
		defer query.Close()
		defer db.Close()
		return (1)
	}
	defer query.Close()
	defer db.Close()
	return (0)
}
