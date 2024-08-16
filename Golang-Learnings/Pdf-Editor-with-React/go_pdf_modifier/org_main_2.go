// Uploading and storing the uploaded file in temp-documents

// package main

// import (
// 	"net/http"
// 	"fmt"
// 	"io/ioutil"
// )

// func uploadDocument(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("File Uplaod Endpoint Hit")

// 	r.ParseMultipartForm(10 << 20)

// 	file, handler, err := r.FormFile("myFile")
// 	if err != nil {
// 		fmt.Println("Error while retrieving the File")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	fmt.Println("Uploaded File: %+v", handler.Filename)
// 	fmt.Println("File Size: %+v", handler.Size)
// 	fmt.Println("MIME Header: %+v", handler.Header)

// 	tempFile, err := ioutil.TempFile("temp-documents", "upload-*.pdf")
// 	if err != nil {
// 		fmt.Println(err)	
// 	}	
// 	defer tempFile.Close()

// 	fileBytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	tempFile.Write(fileBytes)
// 	fmt.Println(w, "Successfully Uploaded File")
// }

// func setupRoutes() {
// 	http.HandleFunc("/upload", uploadDocument)
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	fmt.Println("hello World")
// 	setupRoutes()
// }



// Uploading, storing and displaying the stored file

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"sync"
)

// A map to store the filenames of the uploaded documents.
var uploadedFiles sync.Map

func uploadDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error while retrieving the file")
		fmt.Println(err)
		http.Error(w, "Error while retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-documents", "upload-*.pdf")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while creating a temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while reading the file", http.StatusInternalServerError)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error while writing the file", http.StatusInternalServerError)
		return
	}

	// Store the uploaded file name in the map.
	uploadedFiles.Store(filepath.Base(tempFile.Name()), handler.Filename)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

// Handler to serve the uploaded PDF files.
func serveFiles(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("myFile")
	if fileName == "" {
		http.Error(w, "File name is missing", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join("temp-documents", fileName)
	http.ServeFile(w, r, filePath)
}

// Handler to list the uploaded files.
func listFiles(w http.ResponseWriter, r *http.Request) {
	var response string
	response += "<html><body><h1>Uploaded Files</h1><ul>"

	uploadedFiles.Range(func(key, value interface{}) bool {
		response += fmt.Sprintf("<li><a href=\"/view?file=%s\">%s</a></li>", key.(string), value.(string))
		return true
	})

	response += "</ul></body></html>"

	fmt.Fprintf(w, response)
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadDocument)
	http.HandleFunc("/view", serveFiles)
	http.HandleFunc("/list", listFiles)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
