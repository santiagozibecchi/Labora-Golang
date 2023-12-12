package main

import (
	"fmt"
	"runtime"
	"sync"
)

// A nivel de packege
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// Primitivos de sincronizacion => bloqueadores de exclusion multiple: wg
	wg.Add(1)

	go foo()

	fmt.Println("Luego de la GO ROUTINE")

	bar()

	// Espera a que todas las go routines terminen!
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
