// Em Go, _variables_ são explicitamente declaradas e usadas pelo
// compilador para, por exemplo, verificar se a chamada a uma função
// está correta.

package main

import (
    "fmt"
    "reflect"
)

func main() {

    // `var` declara 1 ou mais variáveis.
    var a = "uma string qualquer"
    fmt.Println(a)

    // Você pode declarar várias variáveis ao mesmo tempo.
    var b, c int = 1, 2
    fmt.Println(b, c) //Imprime na mesma linha o 1 e o 2

    // Go deduzirá os tipos de variáveis inicializadas.
    var d = true
    fmt.Println(d)

    // Variáveis numéricas declaradas sem uma inicialização
    // ela terá valor zero. Por exemplo, na
    // declaração abaixo e ficará com o valor zero
    var e int
    fmt.Println(e)

    // A sintaxe `:=` é uma abreviação para declarar e
    // inicializar uma variável.
    var num01 int = 10
    //O := é para declaração e atribuição de uma só vez
    // O = é só atribuição
    num02 := 5.75
    f := "banana"
    fmt.Println(f)
    fmt.Println(reflect.TypeOf(f))
    fmt.Println(num02)
    fmt.Println(reflect.TypeOf(num02))
    fmt.Println(num01)
    fmt.Println(reflect.TypeOf(num01))
}  