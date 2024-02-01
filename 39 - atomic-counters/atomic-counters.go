// package main

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// )

// func main() {
// 	// Vamos usar um tipo primitivo uint64 para representar nosso
// 	// contador (sempre positivo) e realizar operações atômicas sobre ele.
// 	var ops uint64

// 	// Um WaitGroup nos ajudará a esperar que todas as goroutines
// 	// terminem seu trabalho.
// 	var wg sync.WaitGroup

// 	// Vamos iniciar 50 goroutines que cada uma incrementa o
// 	// contador exatamente 1000 vezes.
// 	for i := 0; i < 50; i++ {
// 		wg.Add(1)

// 		go func() {
// 			for c := 0; c < 1000; c++ {
// 				// Para incrementar atomicamente o contador usamos `AddUint64`.
// 				atomic.AddUint64(&ops, 1)
// 			}

// 			wg.Done()
// 		}()
// 	}

// 	// Esperar até que todas as goroutines estejam concluídas.
// 	wg.Wait()

// 	// Aqui nenhuma goroutine está escrevendo em 'ops', mas usando
// 	// `LoadUint64` é seguro ler atomicamente um valor mesmo enquanto
// 	// outras goroutines estão atualizando-o (atomicamente).
// 	fmt.Println("ops:", atomic.LoadUint64(&ops))
// }

//##########################################################################

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync/atomic"
// )

// var (
// 	requestCount uint64
// )

// func requestHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		atomic.AddUint64(&requestCount, 1)
// 	}
// 	fmt.Fprintf(w, "Request number: %d", atomic.LoadUint64(&requestCount))
// }

// func main() {
// 	http.HandleFunc("/", requestHandler)
// 	fmt.Println("Servidor rodando na porta 8080...")
// 	http.ListenAndServe(":8080", nil)
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
// 	}
// }
//##########################################################################
package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var (
	// Contadores atômicos para cada opção de voto
	votesOptionA uint64
	votesOptionB uint64
)

func voteHandler(w http.ResponseWriter, r *http.Request) {
	option := r.URL.Query().Get("option")

	switch option {
	case "A":
		atomic.AddUint64(&votesOptionA, 1)
		fmt.Fprintf(w, "Voto para a opção A registrado. Total: %d\n", atomic.LoadUint64(&votesOptionA))
	case "B":
		atomic.AddUint64(&votesOptionB, 1)
		fmt.Fprintf(w, "Voto para a opção B registrado. Total: %d\n", atomic.LoadUint64(&votesOptionB))
	default:
		http.Error(w, "Opção inválida", http.StatusBadRequest)
	}
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Resultados da Votação:\nOpção A: %d votos\nOpção B: %d votos\n",
		atomic.LoadUint64(&votesOptionA),
		atomic.LoadUint64(&votesOptionB))
}

func main() {
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/results", resultsHandler)
	fmt.Println("Servidor de votação iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}
