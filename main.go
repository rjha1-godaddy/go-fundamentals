package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}

	fruits = append(fruits, "Orange", "Grape")

	fmt.Println("Fruits:", fruits)

	fruits[1] = "Pineapple"

	for i, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", i, fruit)
	}

	subset := fruits[1:4]
	fmt.Println("Subset:", subset)
}
