/***********
Create by Hugo Janasik
Intern Developer
VMware
************/

package main

import (
	"io"
	"mime/multipart"
	"os"
)

// CreateDir = create image directory if not exist
func CreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

// CreateFolderUpload = create a folder where upload image are store and enable to convert them into StringBase64
func CreateFolderUpload(file multipart.File, header *multipart.FileHeader) {
	out, err := os.Create("./image/" + header.Filename)
	checkError(err)
	defer out.Close()
	_, err = io.Copy(out, file)
	checkError(err)
}
