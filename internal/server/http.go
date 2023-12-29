package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	instance := &HttpServer{}
	go func() {
		initializeController()
		http.ListenAndServe(":8090", nil)
	}()
	return instance
}

func initializeController() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)

		type Msg struct {
			Id   int64  `json:"id"`
			Data string `json:"data"`
		}

		now := time.Now().UnixMilli()

		res := Msg{
			Id:   now,
			Data: "hello",
		}

		fmt.Println(now)
		err := json.NewEncoder(writer).Encode(&res)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	})
}
