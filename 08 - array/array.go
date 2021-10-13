package main
import "fmt"
func main() {
    var a [5]int
    fmt.Println("Array a: ", a)
    a[4] = 100
    fmt.Println("Array a depois de receber 100 na quarta posição: ", a)
    fmt.Println("O elemento na quarta posição: ", a[4])
    fmt.Println("Tamanho do array a: ", len(a))
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array b: ", b)
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("Array 2d: ", twoD)
}