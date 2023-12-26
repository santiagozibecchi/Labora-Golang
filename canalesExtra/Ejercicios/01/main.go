package main

import "fmt"

func main()  {
	
	c := make(chan int)

	// mandamos en valor del canal por medio de una gorutina
	go func ()  {
		c <- 42 
	}()

	// recibe el valor de canal
	fmt.Println(<-c)
}