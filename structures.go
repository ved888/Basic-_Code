package main

import "fmt"

//structure of book type
type Book struct {
	name   string
	author string
}

type Address struct {
	city    string
	country string
}

//nested structure i.e. a structure within the structure
type User struct {
	name    string
	age     int
	address Address //its a nested structure
}

func main() {

	//initialising the book type structure
	b := Book{"Wings of fire", "Dr. A.P.J. Abdul Kalam"}

	fmt.Printf("Book name %s author is %s \n", b.name, b.author)

	//intialising the nested structure

	u := User{
		name: "Amar singh rathour",
		age:  26,
		address: Address{
			city:    "Lamhi",
			country: "India",
		},
	}

	//accessing one by one
	fmt.Println("Name:", u.name)
	fmt.Println("Age:", u.age)
	fmt.Println("City:", u.address.city)
	fmt.Println("Country:", u.address.country)

}
