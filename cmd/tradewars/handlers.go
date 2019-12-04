package main

import (
    "net/http"
    "html/template"
    "log"
    //"fmt"
    "time"
    "github.com/gorilla/websocket"
)

//Message struct for the chat server
type Message struct {
    Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func playersHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("inside playerHandler")
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", 405)
        return
    }

    //fmt.Println("r.Method: ", r.Method)
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
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/players", http.StatusSeeOther) 
        return
    }
    callsign := cookie.Value
    data := struct {
        Callsign string
    }{
        Callsign: "Welcome "+callsign+"!",
    }

	ts, err := template.ParseFiles("./ui/web/navigationscreen.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
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
    var cookie, err = r.Cookie("callsign")
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/players", http.StatusSeeOther) 
        return
    }
    callsign := cookie.Value
    data := struct {
        Callsign string
    }{
        Callsign: callsign,
    }

    ts, err := template.ParseFiles("./ui/web/chatscreen.html")
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.Execute(w, data)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }

    connectionHandler(w, r)

}

func connectionHandler(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
    }

}

func messageHandler() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        var cookie, err = r.Cookie("callsign")
        if err == nil {
            log.Println(cookie)
            http.Redirect(w, r, "/map", http.StatusSeeOther) 
            return
        }
        http.Redirect(w, r, "/players", 303)
    } else {
        return
    }
}