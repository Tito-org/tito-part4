/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"net/http"
	"fmt"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


func method(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method = %s\n", r.Method)
	if r.Method == "POST" {
		resetMysql()
	}
}
