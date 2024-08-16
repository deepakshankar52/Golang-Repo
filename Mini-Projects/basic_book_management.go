package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func bookInfo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getBook(w, r)
	case "POST":
		createBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookUpdateInfo(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
        updateBook(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the 'id' query parameter
	idParam := r.URL.Query().Get("id")
	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		// Find the book with the specified ID
		for _, book := range books {
			if book.ID == id {
				// // If a book viewed once, then decrement its quantity
				// books[i].Quantity--
				// // If quantity is 0, then remove the book from slice
				// if books[i].Quantity == 0 {
				// 	books = append(books[:i], books[i+1:]...)
				// }
				
				json.NewEncoder(w).Encode(book)
				return
			}
		}

		// If the book is not found, return a 404 error
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Get the 'title' query parameter
	titleParam := r.URL.Query().Get("title")

	// Find the book with the specified Title
	if titleParam != "" {
		for _, book := range books {
			if book.Title == titleParam {
				json.NewEncoder(w).Encode(book)
				return
			}
		}

		// If the book is not found, return a 404 error
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// If no ID and Title parameter is provided, return all books
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_book Book
	_ = json.NewDecoder(r.Body).Decode(&new_book)
	books = append(books, new_book)
	json.NewEncoder(w).Encode(new_book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID in URL", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)

	for i, book := range books {
		if book.ID == id {
			books[i].Title = updatedBook.Title
			books[i].Author = updatedBook.Author
			books[i].Quantity = updatedBook.Quantity
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	// If specified ID is not provided, return a 404 error
	http.Error(w, "Item not found", http.StatusNotFound)
}

var books []Book

func main() {
	books = append(books, Book{ID: 1, Title: "Kadavul", Author: "Sujatha", Quantity: 5})
	books = append(books, Book{ID: 2, Title: "Thannir Desam", Author: "Vairamuthu", Quantity: 6})

	// http.HandleFunc("/", Home)
	http.HandleFunc("/books", bookInfo)
	http.HandleFunc("/books/", bookUpdateInfo)	

	fmt.Println("Starting on port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
