package main

import "fmt"

/*
Ponga 100 números en un canal
Extraiga los números del canal y los imprima
*/

func main() {

	c := make(chan int)

	go func () {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	imprimirValores(c)

}

func imprimirValores(c <-chan int) {
	for v := range c {
		fmt.Println("Valor de canal: ", v)
	}
}

