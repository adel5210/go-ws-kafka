package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func initializeWebsocket() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			fmt.Println("Fail to upgrade http connection to ws")
			return
		}
		defer conn.Close()
		fmt.Println("Client connected")

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Fail to read client message")
				return
			}

			fmt.Printf("MESSAGE: %s\n", message)
			conn.WriteMessage(messageType, message)

			if err != nil {
				fmt.Println("Fail to echo client message")
				return
			}
		}
	})
}
