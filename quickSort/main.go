package main

import "fmt"

type arrayType []int

func quickSort(arr arrayType) arrayType {
	if len(arr) <= 2 {
		// array with 2 elements is already sorted
		return arr
	} else {
		//pivot is middle element of array #this is bestcase of this algorithm
		pivot := arr[len(arr)/2]
		var less, greater, equal arrayType

		for _, value := range arr {
			if value < pivot {
				less = append(less, value)
			} else if value > pivot {
				greater = append(greater, value)
			} else {
				equal = append(equal, value)
			}
		}

		return append(append(quickSort(less), equal...), quickSort(greater)...)
	}
}

func main() {
	arr := arrayType{5, 3, 30, 1, 19, 2}
	sorted := quickSort(arr)
	fmt.Println(sorted)
}
