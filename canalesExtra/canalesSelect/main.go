/*
Si tenemos varios canales y estos tienen valores listos para
ser extraidos, podemos utilizar select para extraer el valor de
cualquier canal que tenga un valor listo para ser extraido
*/

package main

import "fmt"

func main()  {
	
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// enviar
	go enviar(par, impar, salir)

	// recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizado!")
}

func enviar(p, i, s chan<- int) {

	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	s <- 0
}


func recibir(p,i,s <-chan int)  {
	// Podemos agarrar valores que esten listos para recibirlos
	for {
		select {
			case v :=  <-p:
				fmt.Println("Desde el canal PAR", v)
			case v :=  <-i:
				fmt.Println("Desde el canal IMPAR", v)
			case v :=  <-s:
				fmt.Println("Salir", v)
				return
		}
	}

}