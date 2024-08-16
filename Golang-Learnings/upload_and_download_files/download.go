package main

import (
	// "fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		// Open the file
		file, err := os.Open("temp-images/temp.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Read the file
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers - used to download the file
		// w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
		// w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Type", "text/plain")


		// Write the file to the response
		w.Write(bytes)
	})

	http.ListenAndServe(":8080", nil)
}