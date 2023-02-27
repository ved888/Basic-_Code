package main

import "fmt"

func main() {
	var a int = 20 // variable declaration
	var ip *int    // pointer variable declaration

	ip = &a /* store address of a in pointer */

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer  */
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	fmt.Printf("Value of a variable before change: %d\n", a)
	//change value using pointer
	*ip = 30
	//access the value using
	fmt.Printf("Value of a variable after change: %d\n", a)

	//example of nil pointer
	println("Example of nil pointer")
	var c *int

	fmt.Printf("The value of varibale c is : %x\n", c)
	fmt.Printf("The address of varibale c is : %x\n", &c)

}
