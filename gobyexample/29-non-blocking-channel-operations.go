package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no messages recieved")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("recived signal", sig)
	default:
		fmt.Println("no activity")
	}
}
