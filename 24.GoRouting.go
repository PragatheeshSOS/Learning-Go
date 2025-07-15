package main

import (
	"fmt"
	"time"
)

func send(ch chan<- string, message string) {
	time.Sleep(5 * time.Second)
	ch <- message
}
func main() {
	fmt.Println("Sending Message..!")
	messageChannel := make(chan string)
	go send(messageChannel, "Completed..!")
	message := <-messageChannel
	fmt.Println("\nReceived: ", message)
}
