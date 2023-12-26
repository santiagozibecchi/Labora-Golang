package main

import "fmt"

/*
	CANAL CON BUFFER
	Reciba la cantidad de elementos o valores del mismo tipo que se pueden quedar en el canal

	Ya no se necesita una gorutina para recibir el valor, lo puedo hacer con una misma gorutina
*/
func main()  {
	// buffered channel
	ca := make(chan int, 2)

	ca <- 42
	ca <- 43

	fmt.Println(<- ca)
	fmt.Println(<- ca)

	fmt.Println("---------------------")

	fmt.Printf("%T\n", ca)

}