// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Criando um temporizador que aguardará 2 segundos
// 	timer := time.NewTimer(2 * time.Second)

// 	// Bloqueando a execução até que o temporizador expire
// 	<-timer.C
// 	fmt.Println("Temporizador Expirou!")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Criando um temporizador de 1 segundo
// 	timer := time.NewTimer(time.Second)

// 	// Executando uma goroutine para aguardar o temporizador expirar
// 	go func() {
// 		<-timer.C
// 		fmt.Println("Temporizador Expirou!")
// 	}()
// 	// Aguardar um pouco para que o temporizador expire
// 	time.Sleep(2 * time.Second)
// 	// Cancelando o temporizador antes que ele expire
// 	if timer.Stop() {
// 		fmt.Println("Temporizador Cancelado!")
// 	} else {
// 		fmt.Println("Temporizador Já Expirou e Não Pode Ser Cancelado!")
// 	}
// }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func tarefaProgramada() {
// 	fmt.Println("Tarefa Executada!")
// }

// func main() {
// 	// Calculando o tempo para a próxima execução da tarefa (daqui a 5 segundos)
// 	duracao := 5 * time.Second
// 	temporizador := time.NewTimer(duracao)

// 	// Bloqueando a execução até que o temporizador expire
// 	<-temporizador.C

// 	// Executando a tarefa programada
// 	tarefaProgramada()
// }
package main

import (
	"fmt"
	"time"
)

func notificacao() {
	fmt.Println("Notificação enviada!")
}

func main() {
	// Criando um temporizador para notificação a cada 10 segundos
	temporizador := time.NewTimer(5 * time.Second)

	// Executando a notificação em um loop infinito
	for {
		<-temporizador.C
		notificacao()

		// Reiniciando o temporizador para o próximo intervalo
		temporizador.Reset(5 * time.Second)
	}
}
