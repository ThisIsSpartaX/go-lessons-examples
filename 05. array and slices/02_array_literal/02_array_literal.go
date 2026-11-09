package main

import "fmt"

func main() {
	coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}
	fmt.Println("Type of coffee:", coffeeTypes)
	fmt.Println("First type of coffee:", coffeeTypes[0])
	fmt.Println("Lenght of the array", len(coffeeTypes))

	coffeeTypes[len(coffeeTypes)-1] = "Milk" // access  last element in the array
	fmt.Println("Type of coffee:", coffeeTypes)

	fmt.Println("String lenght is", len("This is a coffee string"))
}
