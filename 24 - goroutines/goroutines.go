// package main

// import (
//     "fmt"
// )

// func hello() {
//     fmt.Println("Hello world goroutine")
// }
// func main() {
//     go hello()
//     fmt.Println("main function")
// }

// package main

// import (
//     "fmt"
//     "time"
// )

// func hello() {
//     fmt.Println("Hello world goroutine")
// }
// func main() {
//     go hello()
//     time.Sleep(1 * time.Second)
//     fmt.Println("main function")
// }

// package main

// import (
//     "fmt"
// )

// func main() {

//     hello("Martin")
//     hello("Lucia")
//     hello("Michal")
//     hello("Jozef")
//     hello("Peter")
// }

// func hello(name string) {

//     fmt.Printf("Hello %s!\n", name)
// }

// package main

// import (
// 	"fmt"
// )

// func hello(name string) {

// 	fmt.Printf("Hello %s!\n", name)
// }

// func main() {

// 	go hello("Martin")
// 	go hello("Lucia")
// 	go hello("Michal")
// 	go hello("Jozef")
// 	go hello("Peter")
// }

package main

import (
	"fmt"
)

func main() {
	go hello("Martin")
	go hello("Lucia")
	go hello("Michal")
	go hello("Jozef")
	go hello("Peter")
	fmt.Scanln()
}

func hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}
