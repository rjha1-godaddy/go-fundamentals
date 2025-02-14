package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	user1 := User{
		Name:  "Alice",
		Age:   20,
		Email: "alice@wonderland.com",
	}

	user2 := User{"Bob", 25, "bob@robert.com"}
	user3 := User{Name: "Charlie", Age: 30, Email: "charlie@chaplin.com"}

	fmt.Println("User 1:", user1)
	fmt.Println("User 2:", user2)
	fmt.Println("User 3:", user3)

	fmt.Println(user1.Name, "is", user1.Age, "years old.")
}
