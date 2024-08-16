// main.go
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "github.com/gorilla/handlers"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  file, header, err := r.FormFile("file")
  if err != nil {
    http.Error(w, "Error retrieving the file", http.StatusBadRequest)
    return
  }
  defer file.Close()

  out, err := os.Create("/tmp/" + header.Filename)
  if err != nil {
    http.Error(w, "Unable to create the file", http.StatusInternalServerError)
    return
  }
  defer out.Close()

  _, err = io.Copy(out, file)
  if err != nil {
    http.Error(w, "Error saving the file", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "File uploaded successfully: %s", header.Filename)
}

func main() {
  // Initialize CORS options
  allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
  allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
  allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

  http.HandleFunc("/upload", uploadFile)
  corsHandler := handlers.CORS(allowedHeaders, allowedMethods, allowedOrigins)

  fmt.Println("Server started at :8080")
  http.ListenAndServe(":8080", corsHandler(http.DefaultServeMux))
}

