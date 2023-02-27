package main

import (
	"fmt"
	"strings"
	"time"
)

func welcomeMessage() {
	fmt.Println("Welcome to Educative!")
}

func main() {
	/*
		array := [4]string{"A", "B", "C", "D"}
		var lenght int = len(array)
		for j := lenght; j >= 0; j-- {
			for i := 0; i < j; i++ {
				fmt.Print(array[i])

			}

			if j < lenght {
				fmt.Print(" ")
			}
			for k := j - 1; k >= 0; k-- {
				fmt.Print(array[k])

			}
			fmt.Println("")
		}
	*/

	go welcomeMessage()

	fmt.Println("Hello World!")
	time.Sleep(time.Second)
	var my string = "Hello world"
	k := 0
	for _, i := range my {
		k++
		fmt.Print(strings.Repeat(string(i), k))
	}

}
