// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	// Criando três goroutines simulando tarefas independentes
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			fmt.Printf("Goroutine %d executando\n", id)
// 		}(i)
// 	}

// 	// Aguardando todas as goroutines concluírem
// 	wg.Wait()
// 	fmt.Println("Todas as goroutines concluídas")
// }
//##############################################################################
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d iniciado\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d concluído\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 5

// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Todas as tarefas concluídas")
// }

//#########################################################################
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"sync"
// )

// func fetchURL(url string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("Erro ao buscar URL %s: %v\n", url, err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Printf("URL %s retornou com status: %s\n", url, resp.Status)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	urls := []string{"https://example.com", "https://google.com", "https://github.com"}

// 	for _, url := range urls {
// 		wg.Add(1)
// 		go fetchURL(url, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Todas as requisições HTTP concluídas")
// }
//#######################################################################
package main

import (
	"fmt"
	"sync"
)

func processDataBatch(batch []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, num := range batch {
		sum += num
	}
	fmt.Printf("Soma dos elementos no lote: %d\n", sum)
}

func main() {
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	batchSize := 2

	for i := 0; i < len(data); i += batchSize {
		wg.Add(1)
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}
		go processDataBatch(data[i:end], &wg)
	}

	wg.Wait()
	fmt.Println("Processamento de dados em lote concluído")
}
