package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


type usuario struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}


func main()  {


	user := usuario{Name: "Santiago", Age: 28}

	userA, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(userA))
	fmt.Println(string(userA))

}