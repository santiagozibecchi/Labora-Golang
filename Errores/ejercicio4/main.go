package main

import (
	"fmt"
	"os"
)

/*
## **Ejercicio 4: Manejo de errores en operaciones de archivos**

Escribe un programa que intente abrir un archivo cuyo nombre y directorio se pasa como argumento
en la línea de comandos. Si el archivo no existe, el programa debe crearlo. Si hay algún otro error
al abrir el archivo, el programa debe informar del error. No olvides cerrar el archivo
correctamente en todos los casos.
*/

func main()  {
	
	args := os.Args // => argumentos del script / go run . [arg1=name] [arg2=path]
	
	defer recoverAplication()
	
	if len(args) != 3 {
		verificationError := map[string]string{
			"incomplete-args": "Debe ingresar el nombre del archivo y especificar el path en los argumentos de la linea de comando",
		}
		
		panic(verificationError["incomplete-args"])
	}
	
	fileName := os.Args[1]
	path := os.Args[2]
		
	file := openOrCreateFile(fileName, path)
	defer file.Close()

	fmt.Println("El archivo se abrió correctamente.")
}

func recoverAplication() {
	rec := recover()
	if rec != nil {
		fmt.Println("Error:", rec)
	}
}

func openOrCreateFile(fileName string, path string) *os.File {
	fullPath := fmt.Sprintf("%s/%s", path, fileName)
	// Opening file
	file, err := os.Open(fullPath)
	if err != nil {
		// Creating if it doesn't exists
		file, err = os.Create(fullPath)
		if err != nil {
			panic(fmt.Sprintf("No se pudo crear el archivo \"%s\" en la ruta \"%s\" (info: %s)", fileName, path, err.Error()))
		}
		fmt.Printf("Se creó el archivo: %s en la ruta %s\n", fileName, path)
	}
	return file
}

