package main

import (
	"fmt"
	"github.com/adel5210/goWebsocketKafka/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Running main app...")
	server.NewHttpServer()

	sigShut()
}

func sigShut() {
	fmt.Println("Async shutdown ...")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	<-done
	fmt.Println("Shutting down ...")

}
