package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var aaray [5]int
	aaray[0] = 10
	fmt.Println("array ", aaray)

	jasonString, err := json.Marshal(aaray)
	if err != nil {
		panic("Got error ")
	}
	fmt.Println(string(jasonString))

	var b = [3]int{1, 2, 3}
	for k, v := range b {
		fmt.Printf("key %d and value is %v \n", k, v)
	}
}
