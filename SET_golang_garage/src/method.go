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
		http.ServeFile(w, r, "../html/err.html")
		return
	}
	checkError(err)
	defer file.Close()
	CreateFolderUpload(file, header)
	imgBackground = ConvertPicture(w, r,"../background/bg.jpg")
	fileDB = ConvertPicture(w, r, "./image/"+header.Filename)
	if checkSizefile == 0 {
		img2html := "<html> <style> .form-img { background-image: url(data:image/jpeg;base64," + imgBackground + "); height: 195px; background-position: center; background-repeat: no-repeat;} </style> <body class=\"form-img\"> <meta http-equiv=\"refresh\" content=\"0\";URL=\"/\"/> </body> </html>"
		w.Write([]byte(fmt.Sprintf(img2html)))
	}
}
