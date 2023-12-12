package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Hacer una suma de los primeros 100 numeros, de forma secuencial y luego de forma concurrente y
	apreciar la diferencia en el tiempo de ejecución.
*/
/*
	Usando la técnica del sleep sort, hagamos un programa que imprima los primeros 10 números enteros en orden ascendiente.
	¿ Es posible hacerlo en orden descendiente?
*/

func sumaComun(number int )int {

	initTimeComun := time.Now()
	var acc int
	for i := 1; i <= number; i++ {
		acc += i
	}
	 
	fmt.Println("sumaComun", time.Since(initTimeComun))
	return acc
}

func sumaConcurrente(array []int, wg *sync.WaitGroup, canal *chan int) {
	var acc int

	defer wg.Done()

	for _, value := range array {
		acc += value 
	}

	*canal <- acc
}


func main()  {
	
	numbertoSum := 1000000
	var numberArray []int
	canal := make(chan int, 1)

	// var totalSumConcurrent int
	var wg sync.WaitGroup

	for i := 0; i < numbertoSum; i++ {
		numberArray = append(numberArray, i + 1)
	}

	suma1 := sumaComun(numbertoSum)

	wg.Add(2)

	go sumaConcurrente(numberArray[:(len(numberArray)/2)+1], &wg, &canal)
	totalSumConcurrent := <- canal
	
	initTime := time.Now()
	go sumaConcurrente(numberArray[(len(numberArray)/2)+1:], &wg, &canal)
	totalSumConcurrent += <- canal
	end := time.Since(initTime)

	// Cerrar el canal después de que ambas goroutines hayan terminado
	wg.Wait()
	close(canal)

	

	// for valor := range canal {
	// 	totalSumConcurrent += valor
	// }

	fmt.Println("Tiempo Suma concurrente: ", end)


	fmt.Println("Total Suma comun", suma1)
	fmt.Println("TOTAL concurrente: ", totalSumConcurrent)

}