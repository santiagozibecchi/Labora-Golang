package main

import (
	"fmt"
)

// extrae los valores del canal usando una declaración select

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)
	close(salir)

	fmt.Println("A punto de finalizar.")
}

func recibir(c <-chan int, salir chan int) {
	
	for {
		select {
		case v:= <-c:
			fmt.Println("Valor del canal: ", v)
		case v:= <-salir:
			fmt.Println("Saliendo del canal: ", v)
			return
		}
	}
}

// send only (salir)
func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	go func ()  {
		for i := 0; i < 10; i++ {
			c <- i
		}
		salir <- 0
		close(c)
	}()


	return c
}
