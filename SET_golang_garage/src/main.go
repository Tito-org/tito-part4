/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

//Include
import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Declare Gloabl variable
var fileDB string
var marque string
var model string
var gazLevel string
var location string
var checkSizefile int
var imgBackground string

//Main function
func main() {
	CreateDir("image")
	http.HandleFunc("/", GetData)

	fmt.Printf("Connection ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
