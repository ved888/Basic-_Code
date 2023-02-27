package main

import "fmt"

func main() {

	//switch statement to print lower case value of numbers in words less than 5 else print value
	var a int = 4

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Println(a)

	}

	//switch statement to check even odd
	var n int = 5

	switch n%2 == 0 {
	case true:
		fmt.Println("even")
	default:
		fmt.Println("odd")

	}

}
