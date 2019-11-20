// Chevelle Boyer
// Nikita Rubocki

package main

import (
    "log"
    "net/http"
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

    log.Println("Starting server on :4000")
    app.listen(process.env.PORT || 4000, function(){
        console.log("Express server listening on port %d in %s mode", this.address().port, app.settings.env);
    });
    // err := http.ListenAndServe(":4000", mux)
    // log.Fatal(err)
}