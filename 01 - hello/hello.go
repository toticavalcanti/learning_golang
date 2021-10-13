package main

import "fmt"

func main() {
	id := 1
	fmt.Printf("Olá Go!\n")
	key := fmt.Sprintf("user:%d", id)
	fmt.Printf(key)
}