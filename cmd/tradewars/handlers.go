package main

import (
    "net/http"
    "html/template"
    "log"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
    	http.NotFound(w, r)
        return
    }

    ts, err := template.ParseFiles("./ui/web/welcome.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific butt..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
        return
    }
	w.Write([]byte("Create a new butt..."))
}