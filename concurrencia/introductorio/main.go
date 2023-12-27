package main

import (
	"fmt"
	"sync"
)

func main() {
	amount := 100
	routineNumber := 2

	fmt.Println("Suma secuencial:", sumaSecuencial(amount))

	c := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < routineNumber; i++ {
		wg.Add(1)
		fmt.Println( i)
		go enviarGorutinas(c, &wg, amount, routineNumber, i)
	}

	// Esperar a que todas las gorutinas terminen
	go func() {
		wg.Wait()
		close(c) // Cerrar el canal después de que todas las gorutinas hayan terminado
	}()

	// Sumar los resultados recibidos del canal
	sum := 0
	for value := range c {
		sum += value
	}

	fmt.Println("Suma concurrente:", sum)
}

func sumaSecuencial(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// receive-only channel
func enviarGorutinas(c chan<- int, wg *sync.WaitGroup, amount int, routineNumber int, id int) {
	defer wg.Done() // Finalizar la gorutina

	// Rango para cada gorutina
	/*		i			n
		1. (0,1,2,3,...,50)x1 => i=0 gorutine
		2. (0,1,2,3,...,50)x2 => i=1 gorutine
	*/
	start := id * (amount / routineNumber)
	end := (id + 1) * (amount / routineNumber)

	sumaParcial := 0

	// Se suma el rango asignado
	for i := start; i < end; i++ {
		sumaParcial += i
	}

	// Acumulacion de la parcial de los rangos hasta terminar. (receive-only channel)
	c <- sumaParcial
}
