// package main

// import (
// 	"fmt"
// 	"log"
// )

// func main() {
// 	var num1, num2 float64
// 	var operator string

// 	// Taking user input for the first number
// 	fmt.Print("Enter first number: ")
// 	_, err := fmt.Scanln(&num1)
// 	if err != nil {
// 		log.Fatal("Invalid input for the first number")
// 	}
	
// 	// Taking user input for the operator
// 	fmt.Print("Enter operator (+, -, *, /): ")
// 	_, err = fmt.Scanln(&operator)
// 	if err != nil {
// 		log.Fatal("Invalid input for the operator")
// 	}
	
// 	// Taking user input for the second number
// 	fmt.Print("Enter second number: ")
// 	_, err = fmt.Scanln(&num2)
// 	if err != nil {
// 		log.Fatal("Invalid input for the second number")
// 	}

// 	// Performing the calculation based on the operator
// 	var result float64
// 	switch operator {
// 	case "+":
// 		result = num1 + num2
// 	case "-":
// 		result = num1 - num2
// 	case "*":
// 		result = num1 * num2
// 	case "/":
// 		if num2 == 0 {
// 			log.Fatal("Error: Division by zero")
// 		}
// 		result = num1 / num2
// 	default:
// 		log.Fatal("Invalid operator")
// 	}

// 	// Printing the result
// 	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
// }




package main 

import (
	"fmt"
)

func main() {
	num1 , num2 float64
	operator string
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter operater (+, -, *, /) : ")
	fmt.Scanln(&operator)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)
}








