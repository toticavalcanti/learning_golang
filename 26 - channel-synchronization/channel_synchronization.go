// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(done chan bool) {
// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")
// 	done <- true
// }

// func main() {
// 	done := make(chan bool, 1)
// 	go task(done)
// 	<-done
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(taskId int, msg chan int) {
// 	for res := range msg {
// 		fmt.Println("Task: ", taskId, "Message: ", res)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	msg := make(chan int)
// 	go task(1, msg)
// 	go task(2, msg)
// 	go task(3, msg)
// 	go task(4, msg)
// 	for i := 0; i < 10; i++ {
// 		msg <- i
// 	}
// 	 fmt.Scanln()
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func task(done chan bool) {
// 	fmt.Print("Task 1 (goroutine) running...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func task2() {
// 	fmt.Println("Task 2 (goroutine)")
// }

// func main() {
// 	done := make(chan bool, 1)
// 	go task(done)
// 	fmt.Println("Im in the main thread!")

// 	if <-done {
// 		go task2()
// 		fmt.Scanln()
// 	}
// }
package main

import (
	"fmt"
	"time"
)

func task(done chan bool) {
	fmt.Print("Task 1 (goroutine) running...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func task2() {
	fmt.Println("Task 2 (goroutine)")
}

func main() {
	done := make(chan bool, 1)
	go task(done)
	fmt.Println("Im in the main thread!")

	if <-done {
		go task2()
		fmt.Scanln()
	}
}
