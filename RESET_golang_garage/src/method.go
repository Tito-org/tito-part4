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
	if ((r.URL.Path != "/reset") && (r.URL.Path != "/") && (r.URL.Path != "/return")) {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}
	if r.Method == "GET" && r.URL.Path == "/"{
		http.ServeFile(w, r, "../php/index.php")
	}
	if ((r.Method == "POST") && (r.URL.Path == "/reset")) {
		reset()
		refresh := "<html> <script> var timer = setTimeout(function() { window.location='http://"+ os.Getenv("tito_ip") +"' }, 0); </script> </html>"
		w.Write([]byte(fmt.Sprintf(refresh)))
	}
	if ((r.Method == "POST") && (r.URL.Path == "/return")) {
		ret := "<html> <script> var timer = setTimeout(function() { window.location='http://"+ os.Getenv("tito_ip") +"' }, 0); </script> </html>"
		w.Write([]byte(fmt.Sprintf(ret)))
	}
}
