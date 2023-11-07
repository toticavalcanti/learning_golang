package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numSensors         = 5
	dataInterval       = 1 * time.Second
	numWorkers         = 3
	simulationDuration = 5 * time.Second // Defina a duração da simulação
)

func main() {
	var wg sync.WaitGroup

	sensorData := make(chan float64, numSensors) // sensorData transmite leituras de sensores para os workers
	done := make(chan struct{})                  // Canal para sinalizar o término da simulação

	// Inicializa o worker pool
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(sensorData, &wg)
	}

	// Simula a coleta de dados de sensores em goroutines separadas
	for i := 0; i < numSensors; i++ {
		wg.Add(1)
		go simulateSensor(i, sensorData, &wg)
	}

	//Aguarde a duração da simulação
	go func() {
		time.Sleep(simulationDuration)
		close(done)
	}()

	// Aguarde a conclusão de todas as goroutines
	go func() {
		wg.Wait()
		close(sensorData)
	}()

	// Espere que ambas as goroutines terminem
	<-done
	fmt.Println("Coleta de dados concluída.")
}

func worker(sensorData <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range sensorData {
		// Processa os dados do sensor aqui (exemplo: salvar em um banco de dados, realizar cálculos, etc.)
		fmt.Printf("Dado do sensor: %.2f\n", data)
	}
}

func simulateSensor(id int, sensorData chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	// Loop infinito
	for {
		data := rand.Float64() * 100.0 // Simula um dado de sensor
		sensorData <- data
		time.Sleep(dataInterval)
	}
}
