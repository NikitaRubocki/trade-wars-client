// Chevelle Boyer
// Nikita Rubocki

package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/players", welcome)
    mux.HandleFunc("/navigation", navigate)
    mux.HandleFunc("/trade", trade)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting server on :8088")
    err := http.ListenAndServe(":8088", mux)
    log.Fatal(err)
}