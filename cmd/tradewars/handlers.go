package main

import (
    "net/http"
    "html/template"
    "log"
    "fmt"
    "time"
)

func playersHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("inside playerHandler")
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    fmt.Println("r.Method: ", r.Method)
    if r.Method == http.MethodGet {
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
    }else if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            log.Println(err.Error())
            http.Error(w, "Internal Server Error", 500)
        }
        callsign := r.Form.Get("callsign")
        cookie := http.Cookie {
            Name: "callsign",
            Value: callsign,
            Expires: time.Now().AddDate(0, 0, 1),
            Path: "/",
        }
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/map", http.StatusSeeOther)  
    }
}

func mapHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("inside mapHandler")
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error: Could not obtain callsign from cookie", 500)
        return
    }
    callsign := cookie.Value
    fmt.Println(callsign)

	ts, err := template.ParseFiles("./ui/web/navigationscreen.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    data := struct {
        Callsign string
    }{
        Callsign: "Welcome "+callsign+"!",
    }
    err = ts.Execute(w, data)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func tradeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", 405)
        return
    }
	w.Write([]byte("Create a new butt..."))
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    return
}

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://trade-wars-client.herokuapp.com/players", 303)
}