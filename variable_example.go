package main

import "fmt"

func main() {

	//example for variable decleration with initialisation

	var a string = "guys"
	fmt.Printf("varibale type %T and value %v, ", a, a)

	//exmple for variable declaration
	var age int
	//lets assign a value
	age = 10
	fmt.Printf("varibale type %T and value %v, ", age, age)

	//example of variable desclaration without type
	var c = 10
	fmt.Printf("varibale type %T and value %v, ", c, c)

	//example of variable declaration without using keyword var

	z := 20

	fmt.Printf("varibale type %T and value %v, ", z, z)
}
