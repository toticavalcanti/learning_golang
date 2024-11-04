package main

import (
	"fmt"
	"os"
)

func processarArquivo(nomeArquivo string) error {
	// Abre o arquivo
	arquivo, err := os.Open(nomeArquivo)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %v", err)
	}
	// Garante que o arquivo será fechado ao final da função
	defer func() {
		fmt.Println("Fechando arquivo...")
		arquivo.Close()
	}()

	// Processa o arquivo
	fmt.Println("Processando arquivo...")
	// Lê o conteúdo do arquivo e exibe
	buffer := make([]byte, 100)
	for {
		n, err := arquivo.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("erro ao ler arquivo: %v", err)
		}
		fmt.Print(string(buffer[:n]))
	}

	return nil
}

func main() {
	err := processarArquivo("exemplo.txt")
	if err != nil {
		fmt.Printf("Ocorreu um erro: %v\n", err)
		return
	}
	fmt.Println("\nArquivo processado com sucesso!")
}
