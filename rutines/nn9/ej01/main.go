package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {

	fmt.Println("Numero de CPU: ", runtime.NumCPU())
	fmt.Println("Numero de Goroutine: ", runtime.NumGoroutine())

	var wg sync.WaitGroup
	
	wg.Add(2)
	go func ()  {
		wg.Done()
		fmt.Println("Primer Gorutina")

	}()

	fmt.Println("Numero de CPU: ", runtime.NumCPU())
	fmt.Println("Numero de Goroutine: ", runtime.NumGoroutine())

	go func ()  {
		wg.Done()
		fmt.Println("Segunda Gorutina")

	}()

	wg.Wait()

	fmt.Println("Numero de CPU: ", runtime.NumCPU())
	fmt.Println("Numero de Goroutine: ", runtime.NumGoroutine())

	fmt.Println("Finalizo el programa!")
}