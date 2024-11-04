package main

import (
	"fmt"
	"time"
)

func main() {

	// Primeiro, vamos olhar para o rate limiting básico. Suponhamos
	// que queremos limitar o processamento de requisições recebidas.
	// Vamos atender essas requisições a partir de um canal
	// com o mesmo nome.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	// Fechando o canal para indicar que não haverá mais
	// envios de dados, evitando deadlock
	close(requests)

	// Este canal `limiter` receberá um valor a cada 200 milissegundos.
	// Este será o regulador no nosso esquema de rate limiting.
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear na recepção do canal `limiter` antes de atender cada
	// requisição, limitamos a taxa para 1 requisição a cada 200 milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Podemos querer permitir picos curtos de requisições no
	// nosso esquema de rate limiting, preservando o
	// limite geral de taxa. Isso pode ser alcançado por
	// meio do buffer do nosso canal de limitação. Este canal `burstyLimiter`
	// permitirá picos de até 3 eventos.
	burstyLimiter := make(chan time.Time, 3)

	// Preenchemos o canal para representar a capacidade de pico permitida.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200 milissegundos, tentaremos adicionar um novo
	// valor ao `burstyLimiter`, até o seu limite de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora simulamos mais 5 requisições recebidas. As primeiras
	// 3 dessas requisições se beneficiarão da capacidade de pico
	// do `burstyLimiter`.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
