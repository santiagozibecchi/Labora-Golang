package main

import "fmt"

func main()  {
	
	c := make(chan int)

	go enviar(c)

	recibir(c)


	fmt.Println("Finalizando!")
}

// canal send online por parametro de una funcion 
func enviar(c chan<- int) {
	c <- 42
}

// canal send online por parametro de una funcion 
func recibir(c <-chan int) {
	fmt.Println(<-c)
}