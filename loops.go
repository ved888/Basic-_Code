package main

import "fmt"

func main() {

	// for loop to print star pattern
	for a := 0; a < 5; a++ {
		fmt.Print("\n")
		for b := 0; b < a; b++ {
			fmt.Print("*")
		}
	}
	fmt.Print("\n")
	// for loop to print numbers from 1 to n
	var n int = 5
	for a := 1; a <= n; a++ {

		fmt.Println(a)

	}

}
