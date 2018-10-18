/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func request(w http.ResponseWriter, r *http.Request) {
	if ((r.URL.Path != "/refresh") && (r.URL.Path != "/")) {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		http.ServeFile(w, r, "../php/index.php")
	}
	if r.Method == "POST" {
		reset()
	}
	if r.URL.Path == "/refresh" {
		refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://"+ os.Getenv("tito_ip") +"' }, 0); </script> </html>"
		w.Write([]byte(fmt.Sprintf(refresh)))
	}
}
