package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando um canal sem bloqueio para atualização de estoque
	stockUpdateCh := make(chan int, 1)

	// Iniciando uma goroutine para processar as atualizações de estoque
	go func() {
		for productId := range stockUpdateCh {
			// Simulando uma atualização de estoque
			fmt.Printf("Atualizando estoque do produto %d...\n", productId)
			time.Sleep(2 * time.Second)
			fmt.Printf("Estoque do produto %d atualizado\n", productId)
		}
	}()

	// Realizando uma venda e atualizando o estoque do produto
	productId := 123
	fmt.Printf("Realizando venda do produto %d...\n", productId)
	time.Sleep(time.Second)

	// Enviando uma mensagem para o canal sem bloqueio de atualização de estoque
	select {
	case stockUpdateCh <- productId:
		fmt.Printf("Mensagem de atualização de estoque para o produto %d enviada com sucesso\n", productId)
	default:
		fmt.Printf("Falha ao enviar mensagem de atualização de estoque para o produto %d\n", productId)
	}

	// Finalizando a venda
	fmt.Printf("Venda do produto %d realizada com sucesso\n", productId)
}
