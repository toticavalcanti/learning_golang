// package main

// import (
// 	"fmt"
// 	"slices"
// )

// func main() {

// 	// Sorting functions are generic, and work for any
// 	// _ordered_ built-in type. For a list of ordered
// 	// types, see [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
// 	strs := []string{"c", "a", "b"}
// 	slices.Sort(strs)
// 	fmt.Println("Strings:", strs)

// 	// An example of sorting `int`s.
// 	ints := []int{7, 2, 4}
// 	slices.Sort(ints)
// 	fmt.Println("Ints: ", ints)

// 	// We can also use the `slices` package to check if
// 	// a slice is already in sorted order.
// 	s := slices.IsSorted(ints)
// 	fmt.Println("Sorted: ", s)
// }

//----------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"slices"
// 	"time"
// )

// func main() {
// 	datas := []time.Time{
// 		time.Date(2023, time.December, 10, 0, 0, 0, 0, time.UTC),
// 		time.Date(2024, time.July, 20, 0, 0, 0, 0, time.UTC),
// 		time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC),
// 		time.Date(2023, time.July, 20, 0, 0, 0, 0, time.UTC),
// 	}

// 	slices.SortFunc(datas, func(a, b time.Time) int {
// 		if a.Equal(b) {
// 			return 0
// 		}
// 		if a.Before(b) {
// 			return -1
// 		}
// 		return 1
// 	})

// 	fmt.Println("Datas ordenadas:", datas)
// }

//-----------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"slices"
// )

// type Pessoa struct {
// 	Nome  string
// 	Idade int
// }

// func main() {
// 	pessoas := []Pessoa{
// 		{"Maria", 30},
// 		{"João", 25},
// 		{"Ana", 28},
// 	}

// 	// Ordenando por idade
// 	slices.SortFunc(pessoas, func(a, b Pessoa) int {
// 		return a.Idade - b.Idade // Retorna negativo se a < b, zero se a == b, positivo se a > b
// 	})
// 	fmt.Println("Pessoas ordenadas por idade:", pessoas)

// 	// Ordenando por nome
// 	slices.SortFunc(pessoas, func(a, b Pessoa) int {
// 		if a.Nome < b.Nome {
// 			return -1
// 		} else if a.Nome > b.Nome {
// 			return 1
// 		}
// 		return 0
// 	})
// 	fmt.Println("Pessoas ordenadas por nome:", pessoas)
// }

// ---------------------------------------------------------------------------------
package main

import (
	"fmt"
	"slices"
)

type Produto struct {
	Nome  string
	Preço float64
}

func main() {
	produtos := []Produto{
		{"Caderno", 7.50},
		{"Caneta", 3.75},
		{"Mochila", 56.90},
	}

	// Ordenando por preço
	slices.SortFunc(produtos, func(a, b Produto) int {
		if a.Preço < b.Preço {
			return -1
		} else if a.Preço > b.Preço {
			return 1
		}
		return 0
	})
	fmt.Println("Produtos ordenados por preço:", produtos)

	// Ordenando por nome
	slices.SortFunc(produtos, func(a, b Produto) int {
		if a.Nome < b.Nome {
			return -1
		} else if a.Nome > b.Nome {
			return 1
		}
		return 0
	})
	fmt.Println("Produtos ordenados por nome:", produtos)
}
