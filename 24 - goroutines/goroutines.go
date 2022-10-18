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
