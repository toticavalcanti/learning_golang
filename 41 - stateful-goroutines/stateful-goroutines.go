// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync/atomic"
// 	"time"
// )

// type readOp struct {
// 	key  int
// 	resp chan int
// }

// type writeOp struct {
// 	key  int
// 	val  int
// 	resp chan bool
// }

// type deleteOp struct {
// 	key  int
// 	resp chan bool
// }

// type readAllOp struct {
// 	resp chan map[int]int
// }

// func main() {
// 	var readOps uint64
// 	var writeOps uint64
// 	var deleteOps uint64

// 	reads := make(chan readOp)
// 	writes := make(chan writeOp)
// 	deletes := make(chan deleteOp)
// 	readAlls := make(chan readAllOp)

// 	// Goroutine que gerencia o estado
// 	go func() {
// 		var state = make(map[int]int)
// 		for {
// 			select {
// 			case read := <-reads:
// 				read.resp <- state[read.key]
// 			case write := <-writes:
// 				state[write.key] = write.val
// 				write.resp <- true
// 			case delOp := <-deletes: // Renomeei 'delete' para 'delOp'
// 				_, exists := state[delOp.key]
// 				if exists {
// 					delete(state, delOp.key) // Uso correto da função delete
// 					delOp.resp <- true
// 				} else {
// 					delOp.resp <- false
// 				}
// 			case readAll := <-readAlls:
// 				stateCopy := make(map[int]int)
// 				for k, v := range state {
// 					stateCopy[k] = v
// 				}
// 				readAll.resp <- stateCopy
// 			}
// 		}
// 	}()

// 	// Simulando operações de leitura, escrita e deleção
// 	for i := 0; i < 100; i++ {
// 		go func() {
// 			read := readOp{
// 				key:  rand.Intn(5),
// 				resp: make(chan int)}
// 			reads <- read
// 			<-read.resp
// 			atomic.AddUint64(&readOps, 1)
// 		}()
// 	}

// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			write := writeOp{
// 				key:  rand.Intn(5),
// 				val:  rand.Intn(100),
// 				resp: make(chan bool)}
// 			writes <- write
// 			<-write.resp
// 			atomic.AddUint64(&writeOps, 1)
// 		}()
// 	}

// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			delete := deleteOp{
// 				key:  rand.Intn(5),
// 				resp: make(chan bool)}
// 			deletes <- delete
// 			<-delete.resp
// 			atomic.AddUint64(&deleteOps, 1)
// 		}()
// 	}

// 	time.Sleep(time.Second)

// 	// Solicitar e imprimir o estado final
// 	readAll := readAllOp{resp: make(chan map[int]int)}
// 	readAlls <- readAll
// 	stateCopy := <-readAll.resp
// 	fmt.Println("Final state:", stateCopy)

// 	// Impressão dos contadores de operações
// 	fmt.Println("Read operations:", atomic.LoadUint64(&readOps))
// 	fmt.Println("Write operations:", atomic.LoadUint64(&writeOps))
// 	fmt.Println("Delete operations:", atomic.LoadUint64(&deleteOps))
// }

//------------------------------------------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	content string
}

type sendMessage struct {
	msg  message
	resp chan bool
}

type getMessage struct {
	resp chan message
}

func messageQueue() (chan sendMessage, chan getMessage) {
	sendChan := make(chan sendMessage)
	getChan := make(chan getMessage)

	go func() {
		var queue []message
		for {
			if len(queue) > 0 {
				select {
				case msg := <-sendChan:
					queue = append(queue, msg.msg)
					msg.resp <- true
				case get := <-getChan:
					msg := queue[0]
					queue = queue[1:]
					get.resp <- msg
				}
			} else {
				msg := <-sendChan
				queue = append(queue, msg.msg)
				msg.resp <- true
			}
		}
	}()

	return sendChan, getChan
}

func main() {
	send, get := messageQueue()

	var wg sync.WaitGroup

	// Produtor de mensagens
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			respChan := make(chan bool) // create a response channel for each message
			send <- sendMessage{msg: message{content: fmt.Sprintf("Mensagem %d", i)}, resp: respChan}
			ok := <-respChan // wait for the send to be acknowledged
			if !ok {
				fmt.Println("Erro ao enviar mensagem")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Consumidor de mensagens
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			respChan := make(chan message) // create a response channel for each get request
			get <- getMessage{resp: respChan}
			msg := <-respChan // wait for the message to be received
			fmt.Println("Recebido:", msg.content)
		}
	}()

	wg.Wait() // wait for both producer and consumer to finish
}
