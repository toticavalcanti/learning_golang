// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	// Tickers use a similar mechanism to timers: a
// 	// channel that is sent values. Here we'll use the
// 	// `select` builtin on the channel to await the
// 	// values as they arrive every 1000ms.
// 	ticker := time.NewTicker(1000 * time.Millisecond)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case t := <-ticker.C:
// 				fmt.Println("Tick at", t)
// 			}
// 		}
// 	}()

// 	// Tickers can be stopped like timers. Once a ticker
// 	// is stopped it won't receive any more values on its
// 	// channel. We'll stop ours after 1600ms.
// 	time.Sleep(6000 * time.Millisecond)
// 	ticker.Stop()
// 	done <- true
// 	fmt.Println("Ticker stopped")
// }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func enviarEmail() {
// 	fmt.Println("Enviando e-mail...")
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Hour)

// 	for range ticker.C {
// 		enviarEmail()
// 	}
// }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func exibirRelogio() {
// 	fmt.Println("Horário atual:", time.Now().Format("15:04:05"))
// }

// func main() {
// 	ticker := time.NewTicker(1 * time.Second)

// 	for range ticker.C {
// 		exibirRelogio()
// 	}
// }
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func atualizarDadosOnline() {
// 	fmt.Println("Atualizando dados online...")
// }

// func main() {
// 	ticker := time.NewTicker(1000 * time.Second)

// 	go func() {
// 		for range ticker.C {
// 			atualizarDadosOnline()
// 		}
// 	}()

// 	// Simula o tempo para observar as atualizações
// 	time.Sleep(5 * time.Minute)
// 	ticker.Stop()
// 	fmt.Println("Atualizações online interrompidas")
// }
package main

import (
	"fmt"
	"time"
)

func lerSensor() {
	fmt.Println("Lendo sensor...")
}

func main() {
	ticker := time.NewTicker(1000 * time.Minute)

	go func() {
		for range ticker.C {
			lerSensor()
		}
	}()

	// Simula o tempo para observar a leitura dos sensores
	time.Sleep(1 * time.Hour)
	ticker.Stop()
	fmt.Println("Leitura de sensores interrompida")
}
