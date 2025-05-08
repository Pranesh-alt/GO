package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter operator (+, -, *, /, %, ^, &, !): ")
	fmt.Scanln(&operator)

	if operator != "!" {
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printResult(num1, num2, result, operator)
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return add(a, b)
	case "-":
		return subtract(a, b)
	case "*":
		return multiply(a, b)
	case "/":
		return divide(a, b)
	case "&":
		return modulus(a, b)
	case "%":
		return percentage(a, b)
	case "^":
		return power(a, b)
	case "!":
		return factorialOp(a)
	default:
		return 0, errors.New("invalid operator")
	}
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func modulus(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("modulus by zero")
	}
	return float64(int(a) % int(b)), nil
}

func percentage(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero in percentage")
	}
	return a * 100 / b, nil
}

func power(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}

func factorialOp(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("factorial only defined for non-negative integers")
	}
	return float64(factorial(int(a))), nil
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func printResult(num1, num2, result float64, operator string) {
	if operator == "!" {
		fmt.Printf("Result: %.0f! = %.0f\n", num1, result)
	} else {
		fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}
}
