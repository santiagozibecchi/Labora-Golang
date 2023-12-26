package main

func main() {
	
	// SEND ONLY    => PRIMERO CANAL Y LUEGO EL OPERADOR
	cs := make(chan<-int)
	// RECEIVE ONLY => PRIMERO OPERADOR Y LUEGO CANAL
	cs := make(<-chan int)
	// BI-DIRECCIONAL
	cs := make(chan int)

}