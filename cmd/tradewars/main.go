// Chevelle Boyer
// Nikita Rubocki

package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
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