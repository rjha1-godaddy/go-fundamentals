package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return dividend / divisor, nil
}

func main() {
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	product := multiply(5, 3)
	fmt.Println("Product:", product)

	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
	}
}
