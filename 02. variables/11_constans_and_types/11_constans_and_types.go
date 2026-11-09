package main

import "fmt"

func main() {
	// untyped constant - value is int
	const rewardPoints = 10

	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints) // int

	totalRewardPoints := 150.3 // float64

	// Adding untyped constant to a typed variable
	// totalRewardPoints = totalRewardPoints + rewardPoints

	fmt.Printf("Default type of total reward points is %T\n", totalRewardPoints)

	fmt.Printf("Updated loyalty points %.2f\n", totalRewardPoints)
}
