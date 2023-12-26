package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
PATRON DE DISEÑO FANOUT

Trabajo que consume recursos de computos:
Se divide el trabajo en porciones y una gorutina va a estar encargada 
de una porcion de esas, es decir se lanzan un grupo de gorutinas y estas 
se encargan de ejecutar parte de ese trabajo y luego se unifica el 
resultado
*/

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func agregar(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- trabajoConsumeTiempo(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
