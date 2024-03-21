// // No exemplo anterior vimos como gerenciar o estado de um contador simples
// // usando operações atômicas. Para estados mais complexos, podemos usar um _mutex_
// // para acessar dados de forma segura através de múltiplas goroutines.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Container armazena um mapa de contadores; como queremos
// // atualizá-lo concorrentemente a partir de múltiplas goroutines, nós
// // adicionamos um `Mutex` para sincronizar o acesso.
// // Note que mutexes não devem ser copiados, então, se essa
// // `struct` for passada adiante, isso deve ser feito por
// // ponteiro.
// type Container struct {
// 	mu       sync.Mutex
// 	counters map[string]int
// }

// func (c *Container) inc(name string) {
// 	// Bloqueie o mutex antes de acessar `counters`; desbloqueie
// 	// no final da função usando uma instrução [defer](defer).
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.counters[name]++
// }

// func main() {
// 	c := Container{
// 		// Observe que o valor zero de um mutex é utilizável como está, então nenhuma
// 		// inicialização é necessária aqui.
// 		counters: map[string]int{"a": 0, "b": 0},
// 	}

// 	var wg sync.WaitGroup

// 	// Esta função incrementa um contador nomeado
// 	// em um loop.
// 	doIncrement := func(name string, n int) {
// 		for i := 0; i < n; i++ {
// 			c.inc(name)
// 		}
// 		wg.Done()
// 	}

// 	// Execute várias goroutines simultaneamente; note
// 	// que todas acessam o mesmo `Container`,
// 	// e duas delas acessam o mesmo contador.
// 	wg.Add(3)
// 	go doIncrement("a", 10000)
// 	go doIncrement("a", 10000)
// 	go doIncrement("b", 10000)

// 	// Aguarde o término das goroutines
// 	wg.Wait()
// 	fmt.Println(c.counters)
// }
//----------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Logger contém um Mutex para sincronizar o acesso ao recurso compartilhado, neste caso, a saída padrão.
// type Logger struct {
// 	mu sync.Mutex
// }

// // Log imprime uma mensagem no console. O acesso ao método é sincronizado usando o Mutex.
// func (l *Logger) Log(msg string) {
// 	l.mu.Lock()
// 	defer l.mu.Unlock()
// 	fmt.Println(time.Now().Format("15:04:05"), msg)
// }

// func main() {
// 	logger := Logger{}
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			logger.Log(fmt.Sprintf("Goroutine %d is logging", id))
// 		}(i)
// 	}

// 	wg.Wait()
// }
//-------------------------------------------------------------------
package main

import (
	"fmt"
	"sync"
)

// Theater representa um teatro com um conjunto de assentos e um Mutex para controlar o acesso concorrente.
type Theater struct {
	mu    sync.Mutex
	seats map[int]bool // true se o assento estiver reservado
}

// NewTheater inicializa um novo teatro com n assentos disponíveis.
func NewTheater(n int) *Theater {
	seats := make(map[int]bool)
	for i := 1; i <= n; i++ {
		seats[i] = false // todos os assentos estão inicialmente disponíveis
	}
	return &Theater{seats: seats}
}

// Reserve tenta reservar um assento. Retorna true se a reserva foi bem-sucedida.
func (t *Theater) Reserve(seatNumber int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.seats[seatNumber] {
		return false // Assento já reservado
	}
	t.seats[seatNumber] = true // Reserva o assento
	return true
}

func main() {
	theater := NewTheater(10) // Um teatro com 10 assentos
	var wg sync.WaitGroup

	// Simula várias pessoas tentando reservar o mesmo assento ao mesmo tempo
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id, seatNumber int) {
			defer wg.Done()
			success := theater.Reserve(seatNumber)
			if success {
				fmt.Printf("Cliente %d reservou o assento %d\n", id, seatNumber)
			} else {
				fmt.Printf("Cliente %d falhou ao tentar reservar o assento %d (já reservado)\n", id, seatNumber)
			}
		}(i, 5) // Todos tentando reservar o assento número 5
	}

	wg.Wait()
}
