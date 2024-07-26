package main

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Year   int    `json:"year"`
}

// In-memory storage
var (
    books  = []Book{}
    nextID = 1
    mu     sync.Mutex
)

// Handler to add a new book
func addBookHandler(w http.ResponseWriter, r *http.Request) {
    var book Book
    // TODO: Decode the request body into the book struct
    // TODO: Assign a unique ID to the book
    // TODO: Add the book to the in-memory storage

    // TODO: Return the added book as a response
}

// Handler to get the list of all books
func getBooksHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Encode and send the list of books as a response
}

func main() {
    http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
        // Route requests to the appropriate handler based on the HTTP method
        if r.Method == http.MethodPost {
            addBookHandler(w, r)
        } else if r.Method == http.MethodGet {
            getBooksHandler(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    // Start the HTTP server on port 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}
