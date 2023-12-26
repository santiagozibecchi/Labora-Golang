package main

import "fmt"

/*
	CANAL SIN BUFFER
*/
func main()  {
	// unbuffered channel
	ca := make(chan int)
	// Necesito una gorutina que transmita a ese canal y otra reciba de ese canal
	go func ()  {
		ca <- 42
	}()

	fmt.Println(<- ca)
}