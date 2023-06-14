// package main

// import "fmt"

// func main() {
// 	// Criando um canal com capacidade para 2 valores
// 	// Iremos iterar sobre 2 valores no canal `queue`
// 	queue := make(chan string, 2)
// 	queue <- "one"
// 	queue <- "two"
// 	close(queue)

// 	// Iterando sobre os valores recebidos do canal usando `for` e `range`
// 	for elem := range queue {
// 		fmt.Println(elem)
// 	}
// }

//

package main

import (
	"fmt"
	"time"
)

func dataProducer(data chan<- int) {
	for i := 1; i <= 5; i++ {
		data <- i
		time.Sleep(time.Second)
	}
	close(data)
}

func main() {
	data := make(chan int)

	go dataProducer(data)

	for value := range data {
		fmt.Println("Received:", value)
	}
}
