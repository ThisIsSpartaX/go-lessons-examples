package main

import "fmt"

func main() {
	name := "Americano"
	var price float32 = 2.99 // default is float64
	ready := true
	orderedCount := 5
	var stockCount int64 = 5000

	fmt.Printf("Type of name %T\n", name)
	fmt.Printf("Type of price %T\n", price)
	fmt.Printf("Type of ready %T\n", ready)
	fmt.Printf("Type of orderedCount %T\n", orderedCount)
	fmt.Printf("Type of stockCount %T\n", stockCount)
}
