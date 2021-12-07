package main

import "fmt"

func zeroval(ival int) {
    ival = 0
    fmt.Println("(ZEROVAL) Guarda o valor em ival:", ival)
    fmt.Println("(ZEROVAL) &ival guarda o endereço de ival:", &ival)
}

func zeroptr(iptr *int) {
    *iptr = 0
    fmt.Println("(ZEROPTR) *iptr guarda o endereço de i:", iptr)
    fmt.Println("(ZEROPTR) &iptr guarda o endereço de iptr:", &iptr)
    fmt.Println("(ZEROPTR) *iptr guarda o valor que iptr aponta:", *iptr)
}

func main() {
    i := 1
    fmt.Println("(MAIN) i inicial:", i)

    zeroval(i)
    fmt.Println("(MAIN) valor de i na main depois de chamar a zeroval:", i)

    zeroptr(&i)
    fmt.Println("(MAIN) valor alterado do i através da zeroptr:", i)

    fmt.Println("(MAIN) endereço de memória de i:", &i)
}