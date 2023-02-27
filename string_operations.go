package main

import (
	"fmt"
	"strings"
)

func main() {

	//string declration
	var a string = "good Morning"

	fmt.Printf("type %T and value %v\n", a, a)

	//Count function to count occurance of a character in string
	fmt.Println("count for o character", strings.Count(a, "o"))

	//Contains function to check presence of sub string in a string
	fmt.Println("presence of sub string go ", strings.Contains(a, "go"))

	//Index function to get index of character in a string
	fmt.Println("index of sub character o ", strings.Index(a, "o"))

}
