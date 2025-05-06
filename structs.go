package main

import "fmt"

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Pranesh"
	pers1.age = 21
	pers1.domain = "Go"
	pers1.salary = 25000

	// Pers2 specification
	pers2.name = "Peacock"
	pers2.age = 19
	pers2.domain = "Javascript"
	pers2.salary = 15000

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Domain: ", pers1.domain)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Domain: ", pers2.domain)
	fmt.Println("Salary: ", pers2.salary)

}

type Person struct {
	name   string
	age    int
	domain string
	salary int
}
