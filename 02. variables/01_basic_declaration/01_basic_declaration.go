package main

import (
	"fmt"
	"strconv"
)

func main() {
	info := "Medium Espresso price is $3.00"
	fmt.Println(info)
	info = "Medium Espresso price is $4.00"
	fmt.Println(info + " now")
	info = "Medium Espresso price is $5.00"
	fmt.Println(info + " now")
	info += " now"
	fmt.Println(info)
	price := 6
	info2 := "Medium Espresso price is $"
	fmt.Println(info2 + strconv.Itoa(price))

	price2 := 6.00
	info3 := "Medium Espresso price is $"
	fmt.Println(info3 + strconv.FormatFloat(price2, 'f', 2, 64))
}
