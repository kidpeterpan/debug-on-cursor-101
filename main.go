package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	dbConn := os.Getenv("DB_CONN")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf("Port: %s\n", port)
	fmt.Printf("DB Connection: %s\n", dbConn)
	fmt.Printf("Secret Key: %s\n", secretKey) // Careful with logging secrets!

	fmt.Println("Go application running...")

	// Example: Add some logic you might want to debug
	result := add(5, 3)
	fmt.Printf("Result of add: %d\n", result)
}

func add(a, b int) int {
	return a + b // A good place for a breakpoint!
}
