package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	CONCEPTOS	
	Un "mutex" es un bloqueo de exclusión mutua. Los mutexes nos permiten bloquear nuestro código 
	para que solo una gorutina pueda acceder a ese bloque de código bloqueado a la vez.
*/

func main()  {
	fmt.Println("Numero de CPU: ", runtime.NumCPU())
	fmt.Println("Numero de Gorutinas: ", runtime.NumGoroutine())
	
	contador := 0

	const gs = 100
	// wg => ADD, DONE, WAIT
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func ()  {
			mu.Lock()
			// READ
			v := contador
			v++
			// time.Sleep(time.Second * 2)
			// es el yield => luego de la lectura
			runtime.Gosched()
			// WRITE
			contador = v
			mu.Unlock()
			// DONE: cada vez que una gorutine finalice su ejecucion => se ejecuta Done(), resta 1 al contador de Add
			wg.Done()
			fmt.Println("Numero de rutinas: ", runtime.NumGoroutine())
		}()
	}
	
	// Quien monitorea el contador es WAIT => quien da la orden de finalizar
	wg.Wait()
	fmt.Println("Contador: ", contador)

}

