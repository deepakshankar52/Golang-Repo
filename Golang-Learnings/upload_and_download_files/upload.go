package main

import (
    "net/http"
    "fmt"
	"io/ioutil"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Println("File Uplaod Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error while retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded File: %+v", handler.Filename)
	fmt.Println("File Size: %+v", handler.Size)
	fmt.Println("MIME Header: %+v", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)	
	}	
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Println(w, "Successfully Uploaded File")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("hello World")
	setupRoutes()
}
