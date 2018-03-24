package main

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var clients []*websocket.Conn

func removeClient(clientIndex int) {
	clients[clientIndex] = clients[len(clients) - 1]
	clients = clients[:len(clients) - 1]
}

const serveDir = "compressed/"

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, serveDir + "index.html")
	})
	http.HandleFunc("/main.js", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, serveDir + "main.js")
	})
	http.HandleFunc("/main.css", func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, serveDir + "main.css")
	})
	http.HandleFunc("/chatSocket", func (w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		clientIndex := len(clients)
		clients = append(clients, c)
		defer removeClient(clientIndex)
		defer c.Close()
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				return
			}
			for i := 0; i < len(clients); i++ {
				err = clients[i].WriteMessage(mt, msg)
				if (err != nil) {
					return
				}
			}
		}
	})
	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}