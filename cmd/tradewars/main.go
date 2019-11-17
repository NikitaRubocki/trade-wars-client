// Chevelle Boyer
// Nikita Rubocki

package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)

    log.Println("Starting server on :8088")
    err := http.ListenAndServe(":8088", mux)
    log.Fatal(err)
}