package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
)

type Item struct {
    ID    string  `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var items []Item

func main() {
    items = append(items, Item{ID: "1", Name: "Item One", Price: 10.00})

    http.HandleFunc("/items", itemsHandler)
    http.HandleFunc("/items/", itemHandler)

    log.Fatal(http.ListenAndServe(":8000", nil))
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getItems(w, r)
    case "POST":
        createItem(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
        updateItem(w, r)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    id := strings.TrimPrefix(r.URL.Path, "/items/")
    var updatedItem Item
    _ = json.NewDecoder(r.Body).Decode(&updatedItem)
    for i, item := range items {
        if item.ID == id {
            items[i].Name = updatedItem.Name
            items[i].Price = updatedItem.Price
            json.NewEncoder(w).Encode(items[i])
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}
