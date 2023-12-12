package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
	CONCEPTOS
	Un "mutex" es un bloqueo de exclusión mutua. Los mutexes nos permiten bloquear nuestro código
	para que solo una gorutina pueda acceder a ese bloque de código bloqueado a la vez.
*/

func main()  {
	fmt.Println("Numero de CPU: ", runtime.NumCPU())
	fmt.Println("Numero de Gorutinas: ", runtime.NumGoroutine())
	
	var contador int64

	const gs = 100
	// wg => ADD, DONE, WAIT
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func ()  {
			
			// READ
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			// WRITE
			fmt.Println("Contador: ", atomic.LoadInt64(&contador))

			wg.Done()
			}()
			fmt.Println("Numero de Gorutinas: ", runtime.NumGoroutine())
	}
	
	// Quien monitorea el contador es WAIT => quien da la orden de finalizar
	wg.Wait()
	fmt.Println("Contador: ", contador)

}

