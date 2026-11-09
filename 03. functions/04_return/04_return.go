package main

import "fmt"

func calcaulteLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func updateTotalPoints(currentPoints int, points int) int {
	return currentPoints + points
}

func main() {
	totalPoints := 120

	newlyEarnedPoints := calcaulteLoyaltyPoints(9.30) // 18
	fmt.Println("Earned points today", newlyEarnedPoints)

	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints) // 120 + 18 = 138

	fmt.Println("Updated loyal points: today", totalPoints)
}
