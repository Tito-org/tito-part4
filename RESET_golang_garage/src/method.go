/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"fmt"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	if ((r.URL.Path != "/refresh") && (r.URL.Path != "/")) {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}
	if r.Method == "GET" {
		http.ServeFile(w, r, "../html/index.html")
	}
	if r.Method == "POST" {
		reset()
	}
	if r.URL.Path == "/refresh" {
		refresh := "<html> <script> var timer = setTimeout(function() { window.location='https://shwrfr.com/tito/' }, 10); </script> </html>"
		w.Write([]byte(fmt.Sprintf(refresh)))
	}
}
