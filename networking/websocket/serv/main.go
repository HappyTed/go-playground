package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket!")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrader.Upgrade error:", err)
		return
	}
	defer ws.Close()

	// помните, что в соединении с WebSocket вы не можете использовать операторы fmt.Fprintf() для отправки данных клиенту WebSocket
	for {
		mt, message, err := ws.ReadMessage() // читаем только так
		if err != nil {
			log.Println("From", r.Host, "read", err)
			break
		}

		if strings.TrimSpace(string(message)) == "STOP" {
			log.Println("Stop ws server")
			os.Exit(0)
		}

		log.Print("Received:", string(message))

		err = ws.WriteMessage(mt, message) // отправляем только так
		if err != nil {
			log.Println("WriteMessage:", err)
			break
		}
	}
}

/*
для тестирования:

	sudo pacman -S websocat
	websocat ws://localhost:1234/ws
*/
func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/ws", http.HandlerFunc(wsHandler))

	log.Println("Listening to TCP port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}

}
