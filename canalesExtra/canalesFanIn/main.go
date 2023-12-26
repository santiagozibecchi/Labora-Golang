package main 

import (
	"fmt"
	"sync"
)

/*
Range bloquea los canales sino lo cerramos

PATRON DE DISEÑO
Agarrar valores de multiples canales y enviarselos a un solo canal (fanIn)

<- x => enviar => chan<-
	 => recibir <-chan

*/

func main() {
	par := make(chan int)
	impar := make(chan int)
	// canal especifico para recibir los valores de los canales anteriores
	fanin := make(chan int)

	go enviar(par, impar)
	go recibir(par, impar, fanin)
	

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

// send channel
func enviar(par, impar chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// envio al canal par
			par <- i
		} else {
			// envio al canal impar
			impar <- i
		}
	}
	close(par)
	close(impar)
}

// receive channel
// <-chan int (canales que reciben los valores)
// chan<- int (Canal de envio)
func recibir(par, impar <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range par {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range impar {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
