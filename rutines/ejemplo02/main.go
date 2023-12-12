package main

import (
	"fmt"
	"time"
)

func main(){

	now := time.Now()

	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(100*time.Millisecond)
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(100*time.Millisecond)
			fmt.Println(i)
		}
	}()

	interval := time.Since(now)
	fmt.Println(interval.String())

	time.Sleep(2*time.Second)

}


func secuencial()  {
	
	now := time.Now()

	for i := 10; i < 20; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(i)
	}



	interval := time.Since(now)
	fmt.Println(interval.String())


}