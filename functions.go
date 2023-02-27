package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200
	var ret int

	// calling a function to get sum of numbers
	ret = addition(a, b)

	fmt.Printf("Sum of a = %d and b = %d is : %d\n", a, b, ret)

	//example of a function returning more than one values
	println("example of a function returning more than one values")
	x, y := swap("Hello", "World")
	fmt.Println("string is Hello World and the output is ", x, y)
}

// function returning the sum of two numbers
func addition(num1, num2 int) int {

	return num1 + num2
}

func swap(x, y string) (string, string) {
	return y, x
}
