package main

import "fmt"

func main() {
	var coffeeSizes [3]string // fixed size - 3 elements. default element values "".
	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	fmt.Println(coffeeSizes)

	coffeeSizes[1] = "Medium"
	coffeeSizes[2] = "Large"
	fmt.Println(coffeeSizes)

	fmt.Println("First element is:", coffeeSizes[0])
	fmt.Println("Second element is:", coffeeSizes[1])
	fmt.Println("Third element is:", coffeeSizes[2])

	coffeeSizes[2] = "Extra Large"
	fmt.Println(coffeeSizes)

	// coffeeSizes[1] = true // compilation error. type is not string
	// coffeeSizes[5] = "Super Extra Large" // invalid argument: index 5 out of bounds [0:3]
}
