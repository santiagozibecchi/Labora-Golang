package main


func main(){

}


// URL PRACTICA: https://www.notion.so/42b519d81813472d9ac34cf49cf29794?v=53c637d8207b48829132a80afb93eb95&p=a13c67bc8721462e9bce702ca66294b8&pm=c
/*
Ejercicio 1:

Se necesita generar dos secuencias de números que sigan la siguiente lógica

1. Números en orden crecientes (del 1 en adelante)
2. Números que sean primos (1,2,3,5,7,9,11…)

Se pide implementar cada secuencia en una struct que siga cumpla con la interfaz:
type IntSequence interface {

Next() int

}

Por último se pide desarrollar un pequeño programa donde se pueda imprimir los primeros 30 números de una de las dos secuencias según una “tirada de dados”.

Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand).
*/

/*
Ejercicio 2:
Ahora se desea que se imprima un mensaje que diga cuál de las dos secuencias se tomó (según la tirada de dados).

TIP Se le puede declarar un método Title() string en la interface IntSequence y hacer que ambas secuencias lo implementen.
*/

/*
Ejercicio 3:

De forma más general se desea tener una secuencia que me dé números que cumplan una condición (predicado).
Dicha condición será implementada en estructuras (o cualquier otro tipo de dato)  que tengan el un método con la siguiente firma `Fulfill(n int) bool`.
 ***Fulfill*** es un término en inglés que comúnmente se emplea para afirmar que algo cumple una condición.
Además de tener tipos de datos que tengan una función para implementar condiciones de:
Números cualquier sea (para secuencias de enteros lineales)
Números primos
Ahora quiero que agreguen un nuevo tipo que tenga una función que implemente la condición de
Números pares (2,4,6…)
*/

/*
Ejercicio 4:

¿Qué pasaría si quiero una secuencia pero en el orden inverso? 🤔. 
¿Acaso es mucho pedir…? Creo que estoy teniendo sed de combinar dos capacidades: por un lado el orden en el que se generan los números: 
creciente o decreciente y, por otro lado, un predicado que me diga si un número cumple con tal condición.

Se desea implementar una estructura Sequence que tenga como campos un generador y un predicado y que tenga un método llamado Next()
que devuelva el próximo número que se genera usando el generador y que cumpla con el predicado.
*/

// Pair programming: Sistema de Almacenamiento de Documentos
// URL: https://www.notion.so/42b519d81813472d9ac34cf49cf29794?v=53c637d8207b48829132a80afb93eb95&p=6b755089e7f0430181ee94236c13e937&pm=c
