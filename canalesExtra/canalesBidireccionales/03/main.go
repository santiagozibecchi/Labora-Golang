package main

import "fmt"

/*
	CANAL CON BUFFER (chan<-)
*/

func main()  {
	// buffered channel
	// chan<- : es un canal donde solo vamos a poder enviar informacion 
	// receive-only channel
	ca := make(<-chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<- ca)
	fmt.Println(<- ca)

	fmt.Println("---------------------")

	fmt.Printf("%T\n", ca)

}