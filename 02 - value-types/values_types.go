// Go tem vários tipos de valor, incluindo strings,
// números inteiros, floats, booleans, etc. Aqui estão alguns
// exemplos básicos.

package main

import "fmt"

func main() {

    // Strings, podem ser concatenadas com '+'.
    fmt.Println("go" + "lang")

    // Inteiros e floats.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Boleanos, com operadores boleanos.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}