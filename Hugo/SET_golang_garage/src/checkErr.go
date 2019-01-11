/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

//Include
import _ "github.com/go-sql-driver/mysql"

//Check errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
