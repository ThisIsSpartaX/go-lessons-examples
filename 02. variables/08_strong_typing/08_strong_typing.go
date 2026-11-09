package main

import "fmt"

func main() {
	// price of one cup of coffee
	price := 4.50 // float64

	// cups sold in aone day
	quantity := 15 // int

	// total
	total := price * float64(quantity) // mistamatch types
	fmt.Printf("Total income during a day %.2f\n", total)

	// total := int(price) * quantity // mistamatch types
	// fmt.Printf("Total income during a day %d\n", total)
}
