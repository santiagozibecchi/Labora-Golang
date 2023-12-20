package main

import (
	"fmt"
	"strconv"
)

/*
## **Ejercicio 3: Conversión de tipos con manejo de errores**

Escribe una función que convierta un string a un número entero (**`int`**). La función debe devolver
un error si el string no puede convertirse a un número. Prueba tu función con diferentes strings,
incluyendo algunos que no sean números. Usa `panic` y `recover` para manejar los errores
*/

func convertStringToInt(input string) (int, error) {
	
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("No es posible convertir a número")
	}


	return number, nil
}

// invocar un panico
func createPanic(s string) int {
	panic("¡Esto es un pánico!")
}

func main()  {

	invocePanic := true
	inputString := "123A"

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("Error que ha pasado!: %v\n", rec)
		}
	}()

		
	if !invocePanic {
		number, err := convertStringToInt(inputString)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("El numero de salida es:", number)
		}
	} else {
		number:= createPanic(inputString)
		fmt.Printf("'%s' convertido a número: %d\n", inputString, number)
	}

}