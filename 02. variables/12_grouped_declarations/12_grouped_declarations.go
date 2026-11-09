package main

import "fmt"

func main() {
	var coffeeType string = "Latte" // camelCase is not available from other packages
	var quantity int = 3            // camelCase is not available from other packages
	var unitPrice float64 = 4.25    // camelCase is not available from other packages

	fmt.Printf("Ordered %d %s priced at $%.2f each\n", quantity, coffeeType, unitPrice)

	var (
		customerName string = "Bogdan" // camelCase is not available from other packages
		tableNumber  int    = 8        // camelCase is not available from other packages
		isReadyToPay bool   = false    // camelCase is not available from other packages
	)

	if isReadyToPay {
		fmt.Printf("Custimer %s at table %d is ready to pay: yes\n", customerName, tableNumber)
	} else {
		fmt.Printf("Custimer %s at table %d is ready to pay: no\n", customerName, tableNumber)
	}

	const (
		SizeSmaill = "S" // PascalCase is available from other packages. Константы можно не использовать
		SizeMedium = "M" // PascalCase is available from other packages. Константы можно не использовать
		SizeLarge  = "L" // PascalCase is available from other packages. Константы можно не использовать
	)
}
