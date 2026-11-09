package main

import "fmt"

func main() {
	const shopName string = "Brew & Beans"

	// shopName = "Latte Palace" // error: cannot assign to shopName (constant of type string)

	fmt.Println("Welcome to", shopName)
}
