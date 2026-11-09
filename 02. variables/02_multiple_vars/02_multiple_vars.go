package main

import (
	"fmt"
	"strconv"
)

var cofeeName string = "Espresso"

func main() {
	// Type inferred
	size := "Small"

	var price float64 = 2.50

	fmt.Println("Medium Espresso price is $3.00")
	fmt.Println(size, cofeeName, "price is $", price)
	fmt.Println(size + " " + cofeeName + " price is $")

	cofeeName = "Espresso"

	size = "Medium"

	price = 3.50

	fmt.Print(size + " " + cofeeName + " price is $" + strconv.FormatFloat(price, 'f', 2, 64))
	fmt.Printf("\n%s %s price is $%f\n", size, cofeeName, price)
	fmt.Printf("%s %s price is $%.2f\n", size, cofeeName, price)
}
