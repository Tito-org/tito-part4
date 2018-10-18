/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"bufio"
	"encoding/base64"
	"net/http"
	"os"
)

func checkSize(size int64, w http.ResponseWriter, r *http.Request) int {
	if size > 60000 {
		http.ServeFile(w, r, "../html/imgheavy.html")
//		Error := "<html> <style> h1 { position:absolute; top: 150px; left: 300px;} .form-img { background-image: url(data:image/jpeg;base64," + imgBackground + "); height: 195px; background-position: center; background-repeat: no-repeat;} </style> <body class=\"form-img\"> <meta http-equiv=\"refresh\" content=\"1.5\";URL=\"/\"/> <h1> Image To Heavy </h1> </body> </html>"
//		w.Write([]byte(fmt.Sprintf(Error)))
		return (84)
	}
	return (0)
}

//ConvertPicture = convert picture into StringBase64
func ConvertPicture(w http.ResponseWriter, r *http.Request, path string) string {
	var size int64

	FileImg, err := os.Open(path)
	checkError(err)
	defer FileImg.Close()
	fInfo, _ := FileImg.Stat()
	size = fInfo.Size()
	checkSizefile = 0
	if path != "../background/bg.jpg" && checkSize(size, w, r) == 84 {
		checkSizefile = 1
	}
	buf := make([]byte, size)
	fReader := bufio.NewReader(FileImg)
	fReader.Read(buf)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return (imgBase64Str)
}
