/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


func method(w http.ResponseWriter, r *http.Request) {
	resetMysql()
}