/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var nbAlea int

func main() {
	http.HandleFunc("/go", check)
	http.HandleFunc("/stop", check)

	fmt.Printf("Ready\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
