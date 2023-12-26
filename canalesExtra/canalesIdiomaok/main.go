package main

import "fmt"

func main(){

	c := make(chan int)


	go func ()  {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	// a partir de aca, el canal se encuentra cerrado
	v, ok = <-c
	fmt.Println(v, ok)

}