package main

import "fmt"

/*
## **Ejercicio 2: Validación de entrada**
Crea una función que acepte un string y devuelva un error si el string está vacío o si es demasiado largo (por ejemplo, más de 100 caracteres).
*/

func main()  {
	

	value, err := validateStringInput("DuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosaDuiahpiuhfuipasdfiupsahfpuiosa")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El string es valido: ", value)
	}

}


func validateStringInput(input string) (string, error) {

	if len(input) == 0 {
		return input, fmt.Errorf("El string no puede estar vacio!")
	}
 
	if len(input) > 100 {
		return input, fmt.Errorf("El string no puede contener más de 100 caracteres!")
	}

	return input, nil
}