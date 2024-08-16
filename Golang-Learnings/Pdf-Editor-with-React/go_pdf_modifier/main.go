package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/gorilla/mux"
    "github.com/pdfcpu/pdfcpu"
    "github.com/pdfcpu/pdfcpu/api"
)

const (
    uploadPath  = "./uploads/"
    downloadPath = "./downloads/"
)

func main() {
    // Create upload and download directories if they don't exist
    err := os.MkdirAll(uploadPath, 0755)
    if err != nil {
        log.Fatalf("Error creating upload directory: %v", err)
    }
    err = os.MkdirAll(downloadPath, 0755)
    if err != nil {
        log.Fatalf("Error creating download directory: %v", err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/", indexHandler).Methods("GET")
    r.HandleFunc("/upload", uploadHandler).Methods("POST")
    r.HandleFunc("/edit/{filename}", editHandler).Methods("GET")
    r.HandleFunc("/save/{filename}", saveHandler).Methods("POST")
    r.HandleFunc("/download/{filename}", downloadHandler).Methods("GET")

    http.Handle("/", r)

    fmt.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Display upload form
    html := `
    <html>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <input type="file" name="pdfFile" accept=".pdf">
            <input type="submit" value="Upload PDF">
        </form>
    </body>
    </html>
    `
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(html))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the uploaded file
    file, handler, err := r.FormFile("pdfFile")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a unique filename
    filename := handler.Filename
    if filename == "" {
        http.Error(w, "Empty file name", http.StatusBadRequest)
        return
    }

    // Save the uploaded file to the uploadPath directory
    destFile, err := os.Create(filepath.Join(uploadPath, filename))
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, file)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }

    // Redirect to the edit page for the uploaded file
    http.Redirect(w, r, "/edit/"+filename, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    // Extract filename from URL
    vars := mux.Vars(r)
    filename := vars["filename"]

    // Display a simple editing interface
    html := `
    <html>
    <body>
        <h2>Edit PDF: ` + filename + `</h2>
        <iframe src="/viewer/` + filename + `" width="800" height="600"></iframe>
    </body>
    </html>
    `
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(html))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    // Extract filename from URL
    vars := mux.Vars(r)
    filename := vars["filename"]

    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusBadRequest)
        return
    }

    // Get PDF content from form
    pdfContent := r.Form.Get("pdfContent")

    // Write updated PDF content to a temporary file
    tempFile := filepath.Join(uploadPath, "temp_"+filename)
    err = api.WriteFile(tempFile, strings.NewReader(pdfContent), true)
    if err != nil {
        http.Error(w, "Error saving updated file", http.StatusInternalServerError)
        return
    }

    // Move the temporary file to the download directory
    downloadFile := filepath.Join(downloadPath, filename)
    err = os.Rename(tempFile, downloadFile)
    if err != nil {
        http.Error(w, "Error saving updated file", http.StatusInternalServerError)
        return
    }

    // Redirect to download page
    http.Redirect(w, r, "/download/"+filename, http.StatusFound)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
    // Extract filename from URL
    vars := mux.Vars(r)
    filename := vars["filename"]

    // Set headers for file download
    w.Header().Set("Content-Disposition", "attachment; filename="+filename)
    w.Header().Set("Content-Type", "application/pdf")

    // Open the file and stream it to the response
    file, err := os.Open(filepath.Join(downloadPath, filename))
    if err != nil {
        http.Error(w, "Error downloading file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    io.Copy(w, file)
}
