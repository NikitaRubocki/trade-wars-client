// Chevelle Boyer
// Nikita Rubocki

package main

//github.com/gorilla/mux

import (
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

    mux := http.NewServeMux()
    mux.HandleFunc("/", redirect)
    mux.HandleFunc("/players", playersHandler)
    mux.HandleFunc("/map", mapHandler)
    mux.HandleFunc("/trade", tradeHandler)
    mux.HandleFunc("/chat", chatHandler)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    port := os.Getenv("PORT")
    err := http.ListenAndServe(":"+port, mux)
    log.Fatal(err)
}