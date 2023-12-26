package main

import "fmt"


func main()  {
	
	// canal que recibe
	c := gen()

	recibir(c)

	fmt.Println("A punto de finalizar")
}

func recibir(c <-chan int) {
	
	// siempre que usamos range tenemos que cerrar el canal
	// Una vez que se extraen todos los valores del range y no se cierra el canal 
	// este sigue queriendo extraer valores y si no existen, este se bloquea
	for v := range c {
		fmt.Println(v)
	}

}

func gen() <-chan int {
	c := make(chan int)

	// Para enviar valores a un canal tenemos que hacer por gorutinas
	// caso contrario bloquea el codigo
	go func ()  {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)	
	}()

	return c
}