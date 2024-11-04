// package main

// import (
// 	"fmt"
// 	"time"
// )

// const numWorkers = 4

// func downloadWorker(id int, urls <-chan string, done chan<- bool) {
// 	for url := range urls {
// 		fmt.Printf("Worker %d iniciou o download de %s\n", id, url)
// 		// Simula o download
// 		time.Sleep(time.Second)
// 		fmt.Printf("Worker %d concluiu o download de %s\n", id, url)
// 	}
// 	done <- true
// }

// func main() {
// 	urls := []string{"url1", "url2", "url3", "url4", "url5"}
// 	jobs := make(chan string, len(urls))
// 	done := make(chan bool, numWorkers)

// 	for i := 1; i <= numWorkers; i++ {
// 		go downloadWorker(i, jobs, done)
// 	}

// 	for _, url := range urls {
// 		jobs <- url
// 	}
// 	close(jobs)
// 	// Aguarda a conclusão de todos os workers
// 	for i := 1; i <= numWorkers; i++ {
// 		<-done
// 	}
// 	close(done)
// }

//##################################################################################//

// package main

// import (
// 	"fmt"
// 	"time"
// )

// const numWorkers = 2

// func generateReportWorker(id int, tasks <-chan int, done chan<- bool) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d iniciou a geração do relatório %d\n", id, task)
// 		// Simula o processamento
// 		time.Sleep(2 * time.Second)
// 		fmt.Printf("Worker %d concluiu a geração do relatório %d\n", id, task)
// 	}
// 	done <- true
// }

// func main() {
// 	numTasks := 6
// 	tasks := make(chan int, numTasks)
// 	done := make(chan bool, numWorkers)

// 	for i := 1; i <= numWorkers; i++ {
// 		go generateReportWorker(i, tasks, done)
// 	}

// 	for task := 1; task <= numTasks; task++ {
// 		tasks <- task
// 	}
// 	close(tasks)
// 	// Aguarda a conclusão de todos os workers
// 	for i := 1; i <= numWorkers; i++ {
// 		<-done
// 	}
// 	close(done)
// }

//##################################################################################//

// package main

// import (
// 	"fmt"
// 	"time"
// )

// const numWorkers = 3

// func processDataWorker(id int, data <-chan int, results chan<- int) {
// 	for item := range data {
// 		fmt.Printf("Worker %d processando item %d\n", id, item)
// 		// Simula o processamento
// 		time.Sleep(time.Second)
// 		results <- item * 2
// 	}
// }

// func main() {
// 	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	numItems := len(data)
// 	dataChannel := make(chan int, numItems)
// 	results := make(chan int, numItems)

// 	for i := 1; i <= numWorkers; i++ {
// 		go processDataWorker(i, dataChannel, results)
// 	}

// 	for _, item := range data {
// 		dataChannel <- item
// 	}
// 	close(dataChannel)

// 	// Coleta os resultados
// 	for i := 1; i <= numItems; i++ {
// 		result := <-results
// 		fmt.Printf("Resultado: %d\n", result)
// 	}
// 	close(results)
// }

//##################################################################################//

package main

import (
	"fmt"
	"net/http"
)

const numWorkers = 5

func fetchURL(workerID int, urls <-chan string, results chan<- string) {
	for url := range urls {
		fmt.Printf("Worker %d buscando %s\n", workerID, url)
		resp, err := http.Get(url)
		if err == nil {
			results <- fmt.Sprintf("%s: %d", url, resp.StatusCode)
		} else {
			results <- fmt.Sprintf("%s: Erro - %s", url, err)
		}
	}
}

func main() {
	urls := []string{"https://www.example.com", "https://www.google.com", "https://www.github.com", "https://www.openai.com"}
	numURLs := len(urls)
	urlsChannel := make(chan string, numURLs)
	results := make(chan string, numURLs)

	for i := 1; i <= numWorkers; i++ {
		go fetchURL(i, urlsChannel, results)
	}

	for _, url := range urls {
		urlsChannel <- url
	}
	close(urlsChannel)

	// Coleta os resultados
	for i := 1; i <= numURLs; i++ {
		result := <-results
		fmt.Println(result)
	}
	close(results)
}
