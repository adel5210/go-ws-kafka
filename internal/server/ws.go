package server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	wsConns = make([]*websocket.Conn, 0)
)

func initializeWebsocket() {
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := websocket.Upgrade(writer, request, nil, 1024, 1024)
		if err != nil {
			fmt.Println("Fail to upgrade http connection to ws, cause: " + err.Error())
			return
		}
		wsConns = append(wsConns, conn)
		fmt.Println("Client connected")

		defer func() {
			wsConns = func() []*websocket.Conn {
				temp := make([]*websocket.Conn, 0)
				for _, wsConn := range wsConns {
					if conn != wsConn {
						temp = append(temp, wsConn)
					}
				}
				return temp
			}()
			conn.Close()
			fmt.Println("Client disconnected")
		}()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Fail to read client message, cause: " + err.Error())
				return
			}

			fmt.Printf("MESSAGE: %s\n", message)
			//conn.WriteMessage(messageType, message)
			for _, wsConn := range wsConns {
				wsConn.WriteMessage(messageType, message)
			}

			if err != nil {
				fmt.Println("Fail to echo client message, cause: " + err.Error())
				return
			}
		}
	})
}
