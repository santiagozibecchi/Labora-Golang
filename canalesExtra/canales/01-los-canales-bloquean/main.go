package main

import "fmt"

/*
	Hay tipos de canales con Buffer y sin Buffer

	Cuando se utiliza sin buffer?
	Solamente los podemos utilizar cuando tenemos una gorutina
	que transmite al canal, por medio de ese canal y hay otra gorutina del
	otro lado que recibe en ese canal
	Es decir, una envia y otra recibe

	Si tenemos un canal sin buffer y no esta una de esas gorutinas, nuestro
	codigo se bloquea.
*/
func main()  {
	// unbuffered channel
	ca := make(chan int)

	ca <- 42
	fmt.Println(<- ca)
}