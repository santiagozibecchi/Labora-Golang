package main

import "fmt"

func main()  {
	
	c := make(chan int)

	go func ()  {
		for i := 0; i < 5; i++ {
			// envio
			c <- i
		}
		// si no lo cerramos, el for sigue intentando extraer valores del canal
		close(c)
	}()

	for v := range c {
		// Lectura con range sobre los elementos que estan en el canal
		fmt.Println(v)
	}

	fmt.Println("Finalizado!")
}

