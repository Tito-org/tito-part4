/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

//Check the size of the file
func checkSize(size int64, w http.ResponseWriter, r *http.Request) int {
	if size > 60000 {
		http.ServeFile(w, r, "../html/imgheavy.html")
		fmt.Printf("Error Image to heavy\n")
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
	if checkSize(size, w, r) == 84 {
		checkSizefile = 1
	}
	buf := make([]byte, size)
	fReader := bufio.NewReader(FileImg)
	fReader.Read(buf)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return (imgBase64Str)
}
