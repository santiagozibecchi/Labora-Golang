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

func ejercicio05(){

	sequence1 := "ATGCGTAT"
	sequence2 := "ATATGCGT"

	if dnaSequence(sequence1, sequence2){
		fmt.Println("Las secuencias coinciden")
	} else {
		fmt.Println("Las secuencias NO coinciden")
	}
}


func main()  {

	ejercicio05()
	
}