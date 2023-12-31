package main

import "fmt"

func ejercicio01() {
	// Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.

	var number int

	fmt.Print("Ingrese un valor numerico: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error al leer el número:", err)
		return
	}

	if number > 0 {
		fmt.Println("El valor ingresado es mayor que cero.")
	} else if number == 0 {
		fmt.Println("El valor ingresado es igual a cero.")
	} else {
		fmt.Println("El valor ingresado es menor que cero.")
	}
}

func ejercicio02() {
	// Escribir un algoritmo que determine si un número es par.
	var number int

	fmt.Print("Ingrese un valor numerico para determinar si es Par o Impar: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Print("Error al leer el número", err)
	}

	if number%2 == 0 {
		fmt.Print("El número ingresa es PAR\n")
	} else if number%2 != 0 {
		fmt.Print("El número ingresa es IMPAR\n")
	}

}

func ejercicio03() {
	// Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.

	var temp1 float64
	var temp2 float64
	var temp3 float64

	var acctemp float64

	fmt.Print("Ingrese la temperatura 1: ")
	fmt.Scan(&temp1)

	fmt.Print("Ingrese la temperatura 2: ")
	fmt.Scan(&temp2)

	fmt.Print("Ingrese la temperatura 3: ")
	fmt.Scan(&temp3)

	acctemp = temp1 + temp2 + temp3

	fmt.Println("La suma de las temperaturas ambientales es:", acctemp)
	fmt.Println("El promedio de las temperaturas ambientales es:", acctemp/3)

}

func ejercicio04() {
	// Redactar un algoritmo para pasar un periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.

	var day float64
	var hours float64
	var seconds float64
	var minutes float64

	var totalSeconds float64

	fmt.Print("Dia: ")
	fmt.Scan(&day)

	fmt.Print("Horas: ")
	fmt.Scan(&hours)

	fmt.Print("Minutos: ")
	fmt.Scan(&minutes)

	fmt.Print("Segundos: ")
	fmt.Scan(&seconds)

	totalSeconds = day*86400 + hours*3600 + minutes*60 + seconds

	fmt.Printf("El periodo expresado en segundos es: %.2f\n", totalSeconds)
}

func ejercicio05() {

	var x int
	fmt.Print("Ingrese un numero: ")
	fmt.Scan(&x)

	fmt.Println(getRange(x))
}

func getRange(x int) (int, int, int, int, int) {

	s1, s2, s3, s4, s5 := 0, 0, 0, 0, 0

	switch {
	case x >= 0 && x <= 50:
		s1 = x
		return s1, s2, s3, s4, s5
	case x > 50 && x <= 100:
		s1 = 50
		s2 = x - 50
		return s1, s2, s3, s4, s5
	case x > 100 && x <= 700:
		s1 = 50
		s2 = 50
		s3 = x - 100
		return s1, s2, s3, s4, s5
	case x > 700 && x <= 1500:
		s1 = 50
		s2 = 50
		s3 = 600
		s4 = x - 700
		return s1, s2, s3, s4, s5
	case x > 1500:
		s1 = 50
		s2 = 50
		s3 = 600
		s4 = 800
		s5 = x - 1500
		return s1, s2, s3, s4, s5
	default:
		return s1, s2, s3, s4, s5
	}
}

func ejercicioGrupal() {
	// Descripción del Problema:
	// Escribe un programa que calcule el área y el perímetro de un rectángulo.
	// El programa debe pedir al usuario que introduzca la longitud y la anchura del rectángulo y luego calcular y mostrar el área y el perímetro.
	// Especificaciones:
	// Solicita al usuario que ingrese la longitud y la anchura del rectángulo.
	// Calcula el área del rectángulo (longitud * anchura).
	// Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
	// Muestra los resultados (área y perímetro) al usuario.

	var longuitud int
	var ancho int

	fmt.Print("Ingrese la longuitud: ")
	fmt.Scan(&longuitud)
	fmt.Print("Ingrese el ancho: ")
	fmt.Scan(&ancho)

	var area int = longuitud * ancho
	var perimetro int = 2 * (longuitud * ancho)

	fmt.Println("El area del rectangulo es: ", area)
	fmt.Println("El area del rectangulo es: ", perimetro)
}

func forValientPeople () {
/*
1. Para valientes: Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.
	1030 segundos son 17 minutos con 10 segundos
	12045 segundos son 3 horas, 20 minutos con 45 segundos
	176520 segundos son 2 días, 1 hora, 2 minutos con 0 segundos.
*/

	var dateInSeconds int 

	fmt.Print("Ingrese la fecha en segundos: ")
	fmt.Scan(&dateInSeconds)

	var seconds int
	var totalMinutes int
	var minutes int
	var totalHours int
	var hours int
	var day int

	seconds = dateInSeconds % 60
	totalMinutes = (dateInSeconds - seconds)/60
	minutes = totalMinutes%60
	totalHours = (totalMinutes - minutes)/60
	hours = totalHours%24
	day = (totalHours - hours)/24

	if (day != 0) {
		fmt.Printf("%d dias - %d horas - %d minutos - %d segundos", day, hours, minutes, seconds)
	} else if hours != 0 {
		fmt.Printf("%d horas - %d minutos - %d segundos", hours, minutes, seconds)
	} else if minutes != 0 {
		fmt.Printf(" %d minutos - %d segundos", minutes, seconds)
	}


}
type exerciseFunc func()
func executeExercise (fn exerciseFunc) {
	
	fn()

	var optSelected int

	fmt.Println("\n¿Desea seguir realizando el mismo ejercicio? Presione 1 para continuar o 0 para volver al menu")

	_, err := fmt.Scan(&optSelected)
	for optSelected > 1 || optSelected < 0 || err != nil {
		fmt.Println("😡🫵  Por favor! Ingrese 1 para continuar o 0 para volver al menu principal")
		fmt.Println("----- El valor", optSelected, "no es valido -----")
		fmt.Scan(&optSelected)
	}

	for optSelected == 1 {
		fn()
		fmt.Println("\n¿Desea seguir realizando el mismo ejercicio? Presione 1 para continuar o 0 para volver al menu")
		fmt.Scan(&optSelected)
	}

	if optSelected == 0 {
		main()
	}

}

func main() {

	var choosingOptions int
	fmt.Println("\nEliga uno de los siguientes ejercicios:")
	fmt.Println("1. Leer un número e informar si es mayor, igual o menos a cero.")
	fmt.Println("2. Determine si un número es par.")
	fmt.Println("3. Calcular e informar la suma y el promedio de dichos valores.")
	fmt.Println("4. Periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.")
	fmt.Println("5. Suma de una serie de números siguiendo las restricciones impuestas.")
	fmt.Println("6. Obtener area y perimetro de un rectangulo.")
	fmt.Println("7. Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.")
	fmt.Println("0. Presione 0 para salir")

	for true {

		fmt.Print("\nIngrese la opcion: ")
		fmt.Scan(&choosingOptions)

		switch choosingOptions {
		case 1:
			executeExercise(ejercicio01)
		case 2:
			executeExercise(ejercicio02)
		case 3:
			executeExercise(ejercicio03)
		case 4:
			executeExercise(ejercicio04)
		case 5:
			executeExercise(ejercicio05)
		case 6:
			executeExercise(ejercicioGrupal)
		case 7:
			executeExercise(forValientPeople)
		case 0:
			fmt.Println("Usted ha salido del programa")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}

}
