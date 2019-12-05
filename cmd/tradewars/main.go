// Chevelle Boyer
// Nikita Rubocki

package main


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
    mux.HandleFunc("/ws", wsHandler)
    // mux.HandleFunc("/newMsg", writeHandler)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    port := os.Getenv("PORT")
    log.Println("Starting server on port :"+port+"...")
    err := http.ListenAndServe(":"+port, mux)
    log.Fatal(err)

    go messageHandler()
}

// strings.Contains(r.Header.Get("Accept"), "json"){
//     log("JSON REQUEST")
//     w.Header().Set("Content-Type", "application/json")
//     w.Write([]byte(...))
// } else {deal with cookie}

// <script src"/assest.js.main.js" type="text/javascipt">