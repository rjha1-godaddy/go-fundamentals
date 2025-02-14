package main

import "fmt"

func main() {
	userScores := map[string]int{
		"Alice": 95,
		"Bob":   85,
		"Eve":   90,
	}

	userScores["Charlie"] = 88

	score, exists := userScores["Bob"]
	if exists {
		fmt.Println("Bob's score:", score)
	} else {
		fmt.Println("Bob not found")
	}

	fmt.Println("All user scores:")
	for user, score := range userScores {
		fmt.Printf("%s: %d\n", user, score)
	}

	delete(userScores, "Eve")
	fmt.Println("After deletion:", userScores)
}
