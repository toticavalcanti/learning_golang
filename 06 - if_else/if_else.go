// Branching with `if` and `else` in Go is
// straight-forward.

package main

import "fmt"

func main() {

    // Here's a basic example.
    if 7%2 == 0 {
        fmt.Println("7 é par")
    } else {
        fmt.Println("7 é ímpar")
    }

    // You can have an `if` statement without an else.
    if 8%4 == 0 {
        fmt.Println("8 é divisível por 4")
    }

    // A statement can precede conditionals; any variables
    // declared in this statement are available in all
    // branches.

    if num := 9; num < 0 {
        fmt.Println(num, "é negativo")
    } else if num < 10{
        fmt.Println(num, "tem 1 dígito")
    } else {
        fmt.Println(num, "tem múltiplos dígitos")
    }
}

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.