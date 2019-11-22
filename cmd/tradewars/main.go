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
    mux.HandleFunc("/players", welcome)
    mux.HandleFunc("/map", starMap)
    mux.HandleFunc("/trade", trade)
    mux.HandleFunc("/chat", chat)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    port, ok := os.LookupEnv("PORT")
    if ok == false {
        port = "4000"
    }
    err := http.ListenAndServe(":"+port, mux)
    log.Fatal(err)
}