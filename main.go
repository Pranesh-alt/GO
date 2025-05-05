package main

import (
	"fmt"
	"math/cmplx"
)

func loop(num int) int {
    sum := 0 
    for i := 0; i < num; i++ { 
        sum += i 
    }
    return sum 
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	arr1 := []interface{}{1, 2, 3, "pranesh"}
	arr2 := [5]interface{}{4,5,6,7,8}
  
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr1)




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


type Person struct{
	name string
	age int
	domain string
	salary int
}
