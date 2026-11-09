package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	// STEP 1
	// Compile Time (code you write): var coffeePrice = 4.50
	// Runtime (what machine sees): [some momory address] = 4.50

	// STEP 2
	// Compile Time (code you write): fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees): fmt.Println([some momory address], [momory value])
	fmt.Println("Memory address of price 4.50", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory address of price 5.00", &coffeePrice)

	var pointerToCoffeePrice *float64 = &coffeePrice

	*pointerToCoffeePrice = 7.50 // change value via pointer

	fmt.Println("Coffee Price:", coffeePrice)

	changeCoffeePrice(pointerToCoffeePrice)

	fmt.Println("Coffee Price:", coffeePrice)

	fmt.Println("Coffee Price:", *pointerToCoffeePrice) //show value via pointer

	coffeePrice = 10.0

	fmt.Println("Coffee Price:", coffeePrice)
}

func changeCoffeePrice(coffeePricePointer *float64) {
	*coffeePricePointer = 9.00
}
