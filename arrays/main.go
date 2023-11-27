package main

import "fmt"


func dnaSequence(str1 string, str2 string) bool {
	// var sequence []rune = []rune{'A', 'T', 'G', 'C', 'G', 'T', 'A', 'T'}
	strToRune1 := []rune(str1) //  'G', 'C', 'G', 'T', 'A', 'T', 'A', 'T' 
	strToRune2 := []rune(str2) //  'G', 'C', 'G', 'T', 'A', 'T', 'A', 'T' 

	match := true
	
	for i := 0; i < len(strToRune1); i++ {
		
		for j := 0; j < len(strToRune2); j++ {	

			if strToRune1[(i+j)%(len(strToRune1))] != strToRune2[(j)] {

				match = false
				break
			}
		}
	}

	return match
}

func deleteIndexArray(array []int, index int) []int {
	var nuevoArray []int
	for i, value := range array {
	  if i == index {
		continue
	  }
	  nuevoArray = append(nuevoArray, value)
	}
	nuevoArray = append(nuevoArray, 0)
	return nuevoArray
  }

func ejercicio05(){

	sequence1 := "ATGCGTAT"
	sequence2 := "ATATGCGT"

	if dnaSequence(sequence1, sequence2){
		fmt.Println("Las secuencias coinciden")
	} else {
		fmt.Println("Las secuencias NO coinciden")
	}
}

func ejercicio02()  {
	/*
	Haga que reciba un arreglo y una posición y “borre” el valor que hay en dicha posición. 
	Por “borrar” entiendase que no se quita el elemento sino se mueven todos los elementos 
	que están a la derecha un pasito a la izquierda, rellenado con el valor por defecto el 
	extremo derecho. O sea…
	*/
	
	var index int
	fmt.Print("Ingrese la posición que quiere borrar: ")
	fmt.Scan(&index)

	var enterArray []int = []int{0,1,2,3,4,5,6,7}

	newArray  := deleteIndexArray(enterArray, index)

	fmt.Println(newArray)
}

func posiblesOwnDividers(enterNumber int) []int {
	accDivider := []int{}
	// divisores posibles enterNumber/2
	for i := 1; i <= enterNumber/2; i++ {
		if enterNumber % i == 0 {
			accDivider = append(accDivider, i)
		}
	}

	return accDivider
}

func ejercicio01()  {
	/*
	Si quisiera hacer una función para obtener los divisores propios de un número natural (ver ejercicio 1 de la práctica funciones) 
	se pueden usar arreglos como tipo de dato de retorno de la función que hace que calcula los divisores propios? 
	*/

	var naturalNumber int
	fmt.Print("Ingrese un numero natural: ")
	fmt.Scan(&naturalNumber)

	fmt.Printf("Los divisores propios de %d son: %d\n", naturalNumber, posiblesOwnDividers(naturalNumber))
}

func convertTextToSpecialCharacter(text string) string {

	convertedText := ""
	
	for _, character := range text {
		
		switch character {
		case 'a':
			convertedText += "4"
		case 'e':
			convertedText += "3"
		case 'i':
			convertedText += "1"
		case 'o':
			convertedText += "0"
		case 'u':
			convertedText +=	"6"
		default:
			convertedText += string(character)
		}

	}

	return convertedText

}

func ejercicio04()  {
	/*
		4. Desarrolle una función que reciba como parámetro una cadena y que reemplace cada vocal por un carácter 
		que la represente simbólicamente. Se puede usar la siguiente tabla de transformación
		‘a’ → ‘4’
		‘e’ → ‘3’
		‘i’ → ‘1’
		‘o’ → ‘0’
		‘u’ → ‘6’ (no hay mejor candidato, pero si quieren usar otro sean bienvenidos)
		Asumiendo que un usuario ingresa “pepa” y quiere transformar su cadena usando la codificación correspondiente, quedaría “p3p4”.
	*/

	var enterText string

	fmt.Print("Ingrese el texto que quiere convertir: ")
	fmt.Scan(&enterText)

	convertedText := convertTextToSpecialCharacter(enterText)

	fmt.Printf("El texto convertido de %s es: %s\n",enterText, convertedText)
}


func main()  {

	// ejercicio01()
	// ejercicio02()
	// ejercicio04()
	// ejercicio05()
}