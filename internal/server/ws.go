package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func initializeWebsocket() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := websocket.Upgrade(writer, request, nil, 1024, 1024)
		if err != nil {
			fmt.Println("Fail to upgrade http connection to ws, cause: " + err.Error())
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
