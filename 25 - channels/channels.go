// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func CalculateValue(values chan int) {
// 	rand.Seed(time.Now().UnixNano())
// 	value := rand.Intn(10)
// 	fmt.Println("Calculated Random Value: {}", value)
// 	values <- value
// }

// func main() {

// 	values := make(chan int) // create unbuffered channel of integers

// 	//deferred the closing of channel until the end of our main() function’s execution
// 	defer close(values)
// 	// start single goroutine passing in our newly created values channel as its paramete
// 	go CalculateValue(values)
// 	// receives a value from values channel.
// 	value := <-values
// 	fmt.Println(value)
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func CalculateValue(c chan int) {
// 	fmt.Println(time.Now().UnixNano())
// 	rand.Seed(time.Now().UnixNano())
// 	value := rand.Intn(10)
// 	fmt.Println("Calculated Random Value: {}", value)
// 	time.Sleep(1000 * time.Millisecond)
// 	c <- value
// 	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
// }

// func main() {

// 	valueChannel := make(chan int, 3)
// 	defer close(valueChannel)

// 	go CalculateValue(valueChannel)
// 	go CalculateValue(valueChannel)

// 	values := <-valueChannel
// 	fmt.Println(values)
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("This executes regardless as the send is now non-blocking")
}

func main() {

	valueChannel := make(chan int, 2)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)

	time.Sleep(1000 * time.Millisecond)
}
