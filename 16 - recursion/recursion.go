package main

import "fmt"

func fact(n int) int {
    if n < 2 {
        //caso base
        return 1
    }
    return n * fact(n-1)
}

func main() {
    //declara uma variável do tipo função anônima(closure)
    //que recebe um parâmetro do tipo inteiro e retorna um inteiro
    var fib func(n int) int
    //usa a variável criada anteriormente para receber o retorno
    //da implementação da função anônima fib
    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    //imprime os resultados
    fmt.Println("O fatorial de 7 é:", fact(7))
    fmt.Println("O sétimo número da sequência de Fibonnaci é:", fib(7))
}