package main

import (
    "net/http"
    "html/template"
    "log"
    "fmt"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        callsign := r.Form.Get("callsign")
        fmt.Println(callsign)
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

func starMap(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/web/navigationscreen.html")
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

func trade(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
        return
    }
	w.Write([]byte("Create a new butt..."))
}

func chat(w http.ResponseWriter, r *http.Request) {
    return
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://127.0.0.1:4000/players", 303)
}