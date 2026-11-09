package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Dessert Menu:", dessertMenu)

	slice := dessertMenu[1:2] // from index 1 to index 2 (not included), or [1 , 2), or 1 >= elements < 2
	fmt.Println("Slice of the dessert Menu [1:2]:", slice)

	slice = dessertMenu[:] // all elements
	fmt.Println("Slice of the dessert Menu [:]:", slice)

	slice = dessertMenu[2:] // all elements from index 2 to end
	fmt.Println("Slice of the dessert Menu [2:]:", slice)

	slice = dessertMenu[:3] // all elements from beginning to index 3 (not included)
	fmt.Println("Slice of the dessert Menu [:3]:", slice)
}
