package main

import (
	"fmt"
)

func main() {

	//example of adding two values
	var a int = 10
	var b int = 20
	//without using printf function formating of output
	fmt.Println("addition of a =", a, " and b = ", b, "is ", a+b)

	//example of calculating the are of circle
	const PI float32 = 22.7

	var r float32 = 4
	var area_of_circle = PI * (r * r)

	fmt.Println("area of circle with radius =", r, "is ", area_of_circle)

}
