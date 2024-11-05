package main

import (
	"fmt"
	"time"
)

func registrarTempoExecucao(operacao string) func() {
	inicio := time.Now()
	return func() {
		duracao := time.Since(inicio)
		fmt.Printf("Operação %s levou %v para executar\n", operacao, duracao)
	}
}

func operacaoLonga() {
	defer registrarTempoExecucao("operacaoLonga")()

	// Simula uma operação que leva tempo
	fmt.Println("Iniciando operação longa...")
	time.Sleep(2 * time.Second)
	fmt.Println("Operação longa finalizada!")
}

func main() {
	operacaoLonga()
}
