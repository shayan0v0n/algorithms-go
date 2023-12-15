package main

import (
	"fmt"
	"math"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func binarySearch(target, start, end int) int {
	middle := int(math.Round((float64(start) + float64(end)) / 2))

	if numbers[middle] == target {
		return middle
	}

	if numbers[middle] < target {
		return binarySearch(target, start+1, end)
	}

	if numbers[middle] > target {
		return binarySearch(target, start, end-1)
	}

	return 0
}

func main() {
	res := binarySearch(5, 0, len(numbers))
	fmt.Printf("your index is %v, and your value is %v", res, numbers[res])
}
