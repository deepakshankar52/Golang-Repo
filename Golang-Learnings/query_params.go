package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    http.HandleFunc("/user", userHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    // Accessing query parameters
    queryParams := r.URL.Query()
    name := queryParams.Get("name")
    age := queryParams.Get("age")

    // Accessing headers
    contentType := r.Header.Get("Content-Type")

    // Print query parameters and headers
    fmt.Println("Query Parameters:")
    fmt.Printf("Name: %s, Age: %s\n", name, age)
    fmt.Println("Headers:")
    fmt.Printf("Content-Type: %s\n", contentType)

    // Reading request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Unmarshalling JSON body
    var user User
    if err := json.Unmarshal(body, &user); err != nil {
        http.Error(w, "Error unmarshalling JSON body", http.StatusBadRequest)
        return
    }

    // Print user details from request body
    fmt.Println("Request Body:")
    fmt.Printf("Name: %s, Age: %d\n", user.Name, user.Age)

    // Sending response
    response := map[string]interface{}{
        "message": "User details received successfully",
        "user":    user,
    }
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error marshalling JSON response", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}
