package main

import (
	"fmt"
	"strings"
)

/*
1. Define una estructura llamada **`Persona`** con los siguientes campos: **`nombre`**, **`edad`**, **`ciudad`** y **`telefono`**.
2. Crea dos variables de tipo **`Persona`** con valores iniciales distintos.
3. Imprime los valores de ambas variables por pantalla.
4. Define una función llamada **`cambiarCiudad`** para cambiar la ciudad de la persona**.** Si la ciudad es la misma, no hace nada.
5. Llama a la función **`cambiarCiudad`** con una de las personas creadas anteriormente y una ciudad distinta a la actual.
*/

type Persona struct {
	nombre string;
	edad int;
	ciudad string;
	telefono string;
}

func (p *Persona) cambiarCuidad(cuidad string) {
	p.ciudad = cuidad
}

func ejercicio01()  {
	p1 := Persona{"Santiago", 28, "Resistencia", "3777541647"}
	p2 := Persona{"Andres", 28, "Corrientes", "377554422"}

	fmt.Println(p1)
	fmt.Println(p2)

	p1.cambiarCuidad("Goya")

	fmt.Println(p1)
}

type persona struct {
	nombre string
	edad int
	email string
	direccion direccion
}

type direccion struct {
	calle string
	numero int
}

func ejercicio02()  {
	
	var persona1  persona

	persona1.nombre = "Santiago"

	direccion1 := direccion{"Calle X", 200}


	fmt.Println(persona1)
	fmt.Println(direccion1)

}

func ejercicio03()  {
	/*
		Defina una estructura de datos para manejar conjuntos de enteros de la siguiente forma.  
			`Add(value int) int` 
			Agrega al final. Se incrementa en uno la longitud.

		`RemoveAt(index int) bool`
		True si lo borró, false si no pudo (no importa el motivo). Si se borra entonces se decrementa en 
		uno la longitud
			`Lenght() int`

		Longitud: Cantidad de elementos que tiene. Esta determinada por la cantidad de valores que 
		se agregan menos aquellos que se borran.
			`Set(value int, index int) bool`

		True si lo modificó, false si no pudo (no importa el motivo)
			`ToString() string`

		Libre elección. La función strings.Join (del paquete strings) puede ser de útil.
		Para pensar: Tanto el `RemoveAt` y el `Set` son métodos cuya ejecución puede “fallar” si 
		se pasa como `index` un valor que esté fuera de los “límites” (dado por la longitud)… 
		Yo pregunto entonces 🤔 ¿Es esto algo entonces que pueda ser codificado en una función 
		aparte de forma tal que se re utilice en ambos métodos?
		`isInBoundaries(index) bool`
	*/

}

func deleteIndexArray(array []int, index int) []int {
	var nuevoArray []int
	for i, value := range array {
	  if i == index {
		continue
	  }
	  nuevoArray = append(nuevoArray, value)
	}
	return nuevoArray
}

type IntegerSet struct {
	value []int
}

func (integerSet *IntegerSet) add(number int) {
	integerSet.value = append(integerSet.value, number)
}

func (integerSet *IntegerSet) removeAt(index int) bool {

	if isInBoundaries(index, len(integerSet.value)) {
		integerSet.value = deleteIndexArray(integerSet.value, index)
		return true
	}
	return false
}

func (integerSet *IntegerSet) Length() int {
	return len(integerSet.value)
}

func isInBoundaries(index int, setLength int) bool {
	return index >= 0 && index < setLength
}

func (integerSet *IntegerSet) ToString() string {
	strValues := make([]string, len(integerSet.value))
	for i, v := range integerSet.value {
		strValues[i] = fmt.Sprint(v)
	}
	return strings.Join(strValues, ", ")
}

func newIntegerSet() *IntegerSet {
	return &IntegerSet{}
}

func main()  {
	
	// ejercicio01()
	// ejercicio02()

	fmt.Println(newIntegerSet().value)


}