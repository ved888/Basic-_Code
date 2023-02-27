package main

import "fmt"

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

//example of empty type interface
func customePrint(i interface{}) {
	fmt.Println(i)
}

func main() {
	//example of interface
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("Amar")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))

	//example of empty type interface
	var myType interface{} //myType valiable is of type interface

	//now assigning value to out myType interface
	myType = 10
	//print using our print mehtod
	customePrint(myType)
	myType = "Good morning India "
	customePrint(myType)
}
