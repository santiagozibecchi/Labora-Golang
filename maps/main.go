package main

import "fmt"

/*
	slice estrucas que referencias a un valor, igual para los mapas
*/

func main() {
/*
	var isPrimeByInt = make(map[int]bool)


	for i := 0; i < 100; i++ {
		if isPrimer(i) {
			isPrimeByInt[i] = true
		}
	}

	fmt.Println(isPrimeByInt)
*/

/*
	var ocurrenciesByChar map[string]int = make(map[string]int)

	for _, char := range "Hola amigos" {
		_, exist := ocurrenciesByChar[string(char)]

		if !exist {
			ocurrenciesByChar[string(char)] = 0
		}
		ocurrenciesByChar[string(char)]++
	}

	fmt.Printf("%+v", ocurrenciesByChar)
*/
	var ocurrenciesByChar map[rune]int = make(map[rune]int)

	for _, char := range "Hola amigos" {
		ocurrenciesByChar[char]++
	}

	fmt.Printf("%+v\n", ocurrenciesByChar)


	// taken frot herpss//80.atk1padla.org/utk1/1986_of_guntsplanetary_systems
	// planetCountBySystem := map[string]int {
	// 	"sol": 9,
	// 	"Luna": 9,
	// }
	
	// planet, exist := planetCountBySystem["sol"]

	// fmt.Println(planet, exist)
}

func isPrimer(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

