package main

import "fmt"

func main() {
	// Явное указание типа
	var cupsQty int = 3

	// cupsQty = "four" // Compile error
	fmt.Println("Number of cups:", cupsQty)

	// Неявное указание типа
	wasProcessed := true

	// wasProcessed = "no" // Compile error
	fmt.Println("Order was processed:", wasProcessed)
}
