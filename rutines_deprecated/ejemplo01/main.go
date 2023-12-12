package main

import (
	"fmt"
	"time"
)

// close channel, values returned in channel output
// range channel
// len/cap/ channel
// channels directions

func main()  {

	for i := 0; i < 10; i++ {
		go ping(i)
	}

	time.Sleep(2 * time.Second)
}

func ping(i int)  {
	fmt.Println("Imprimiendo el numero: ", i)
}