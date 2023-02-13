// package main

// import (
// 	"fmt"
// 	"time"
// )

// func server1(ch chan string) {
// 	// sleeps for 6 seconds then writes the
// 	// text "from server1" to the channel ch
// 	time.Sleep(6 * time.Second)
// 	ch <- "from server1"
// }
// func server2(ch chan string) {
// 	// sleeps for 3 seconds then writes the
// 	// text "from server2" to the channel ch
// 	time.Sleep(3 * time.Second)
// 	ch <- "from server2"

// }
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)
// 	// calls the go Goroutines server1
// 	go server1(output1)
// 	// calls the go Goroutines server2
// 	go server2(output2)
// 	// select statement blocks until one
// 	// of its cases is ready
// 	select {
// 	case s1 := <-output1:
// 		fmt.Println(s1)
// 	case s2 := <-output2:
// 		fmt.Println(s2)
// 	}
// }

// ----------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func process(ch chan string) {
// 	// sleeps for 10.5 seconds then writes the
// 	// text "process successful" to the channel ch
// 	time.Sleep(10500 * time.Millisecond)
// 	ch <- "process successful"
// }

// func main() {
// 	ch := make(chan string)
// 	go process(ch)
// 	//infinite loop
// 	for {
// 		// sleeps for 1000 milliseconds (1 second)
// 		// during the start of each iteration
// 		time.Sleep(1000 * time.Millisecond)
// 		// and then performs a select operation
// 		select {
// 		case v := <-ch:
// 			// will be executed after 10.5 seconds when func
// 			// process will write in channel "process successful"
// 			fmt.Println("received value: ", v)
// 			return
// 		default:
// 			fmt.Println("no value received")
// 		}
// 	}
// }

// ----------------------------------------------------------------

// package main

// func main() {
// 	ch := make(chan string)
// 	select {
// 	case <-ch:
// 	}
// }

// ----------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	ch := make(chan string)
// 	select {
// 	case <-ch:
// 	default:
// 		fmt.Println("default case executed")
// 	}
// }

// ------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func server1(ch chan string) {
// 	ch <- "from server1"
// }
// func server2(ch chan string) {
// 	ch <- "from server2"

// }
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)
// 	go server1(output1)
// 	go server2(output2)
// 	time.Sleep(1 * time.Second)
// 	select {
// 	case s1 := <-output1:
// 		fmt.Println(s1)
// 	case s2 := <-output2:
// 		fmt.Println(s2)
// 	}
// }

// ---------------------------------------------------

package main

func main() {
	select {}
}
