package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Trabalho recebido:", j)
			} else {
				fmt.Println("Todos os trabalhos recebidos")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Trabalho enviado:", j)
	}
	close(jobs)
	fmt.Println("Todos os trabalhos enviados")

	<-done
}
