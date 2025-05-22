package main

import "fmt"

// kidsWithCandies determines which kids can have the greatest number of candies
// if they are given all the extraCandies.
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}            // Stores the final boolean results
	maxCandies := getMax(candies) // Find the current maximum in the list

	for _, currentCandies := range candies {
		if currentCandies+extraCandies >= maxCandies {
			result = append(result, true) // This kid can reach or exceed max
		} else {
			result = append(result, false)
		}
	}

	return result
}

// getMax returns the maximum value from an integer slice
func getMax(arr []int) int {
	max := arr[0] // Initialize with the first element
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	// Example test case
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies)) // Output: [true true true false true]
}
