package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i = i + 2
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println("nextInt chamada", nextInt())
    fmt.Println("nextInt chamada", nextInt())
    fmt.Println("nextInt chamada", nextInt())
	fmt.Println()
    newInts := intSeq()
    fmt.Println("newInts chamada", newInts())
}