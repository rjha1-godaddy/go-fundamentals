package main

import "fmt"

func main() {
	// If-Else Statement
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	fmt.Println()

	// For Loop
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// While Loop
	fmt.Println("Counting from 5 to 1:")
	i := 5
	for i >= 1 {
		fmt.Println(i)
		i--
	}

	fmt.Println()

	// Nested Loop
	fmt.Println("Multiplication Table:")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	fmt.Println()

	// Break Statement
	fmt.Println("Counting from 1 to 10:")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println()

	// Continue Statement
	fmt.Println("Counting from 1 to 10 (skipping 6):")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println()

	// Switch Case
	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week!")
	case "Friday":
		fmt.Println("Weekend eve!")
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hang in there!")
	}
}
