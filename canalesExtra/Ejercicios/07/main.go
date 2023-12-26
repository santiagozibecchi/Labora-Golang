/*
Lance 10 gorutinas
Cada gorutina agrega 10 números a un canal
Extrae los números del canal e imprímelos
*/
/*
Lance 10 gorutinas
Cada gorutina agrega 10 números a un canal
Extrae los números del canal e imprímelos
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup

	// Use WaitGroup para esperar que todas las gorutinas finalicen
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go enviarGorutinas(c, &wg)
	}

	// Esparar a que todas las gorutinas terminen
	go func() {
		wg.Wait()
	}()

	imprimirValores(c)
}

// WaitGroup como un puntero para que todas las gorutinas compartan la misma instancia
func enviarGorutinas(c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Finalizar la gorutina
	for i := 0; i < 10; i++ {
		c <- i
	}
	fmt.Println("--------")
}

func imprimirValores(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}
