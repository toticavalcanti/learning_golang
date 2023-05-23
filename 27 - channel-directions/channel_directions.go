package main

import "fmt"

func ping(pings chan<- string, msg string) {
	//put the received message in the pings channel output
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	//read the message that was passed to pings channel and assign to msg
	msg := <-pings
	//put the msg content in the pongs channel output
	pongs <- msg
	//pings <- "Trying send message to pings channel"
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	//call ping function passing pings channel to it
	ping(pings, "Sending message to the pings channel")
	//Read message from pings channel inside pong function
	pong(pings, pongs)
	//print the output from pongs channel
	fmt.Println(<-pongs)
}
