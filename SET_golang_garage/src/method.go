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

//methodGet = Manage query http GET
func methodGet(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../html/index.html")
}

//methodPost = Manage query http POST
func methodPost(w http.ResponseWriter, r *http.Request) {
	model = r.PostFormValue("model")
	marque = r.PostFormValue("marque")
	gazLevel = r.PostFormValue("gazLevel")
	location = r.PostFormValue("location")
	file, header, err := r.FormFile("file")
	if marque == "" || model == "" || gazLevel == "" || location == "" || file == nil {
		fmt.Printf("marque = '%s'\nmodel = '%s'\ngazLevel = '%s'\nlocation = '%s'\nfile = '%s'\n", marque, model, gazLevel, location, file)
		http.ServeFile(w, r, "../html/err.html")
		fmt.Printf("Error complete correctly the form\n")
		return
	}
	checkError(err)
	defer file.Close()
	CreateFolderUpload(file, header)
	fileDB = ConvertPicture(w, r, "./image/"+header.Filename)
	if checkSizefile == 0 {
		img2html := "<html> <style> .form-img { background-image: url(\"http://www.wallpapermaiden.com/wallpaper/21181/download/2560x1600/long-road-field-clean-sky-snow-mountains.jpg\"); background-position: center; background-repeat: no-repeat; background-attachment: fixed; -webkit-background-size: cover; -moz-background-size: cover; -o-background-size: cover; background-size: cover;} </style> </html>"
		refresh := "<html> <body class=\"form-img\"> <meta http-equiv=\"refresh\" content=\"0\";URL=\"/\"/> </body> </html>"
		w.Write([]byte(fmt.Sprintf(img2html)))
		w.Write([]byte(fmt.Sprintf(refresh)))
	}
}
