package main

import "fmt"

func createTemperatureAjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	ajustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}

	return ajustTemperature, baseTemperature
}

func main() {
	// _, originalTemp := createTemperatureAjuster() // Если не используем один возвращаемый параметр
	ajustTemp, originalTemp := createTemperatureAjuster()

	// 90.0
	fmt.Printf("Oirignal Temperature is %.1f\n", originalTemp)

	// +1.5 -> 91.5
	fmt.Printf("Ajusted Temp is +1.5:%.1f\n", ajustTemp(1.5))

	// -3.0 -> 88.5
	fmt.Printf("Ajusted Temp is -3.0:%.1f\n", ajustTemp(-3.0))

	// +5.0 -> 93.5
	fmt.Printf("Ajusted Temp is +5.0:%.1f\n", ajustTemp(5.0))

	// 90
	fmt.Printf("Oirignal Temperature is %.1f\n", originalTemp) // not changed
}
