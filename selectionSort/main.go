package main

import "fmt"

func deleteElement(arr []int, i int) (newArr []int, val int) {
	val = arr[i]
	newArr = append(arr[:i], arr[i+1:]...)
	return
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func sortFunc(oldArr []int) []int {
	sortedArr := []int{}
	arrLen := len(oldArr)
	for i := 0; i < arrLen; i++ {
		smallest := findSmallest(oldArr)
		newArr, val := deleteElement(oldArr, smallest)
		oldArr = newArr
		sortedArr = append(sortedArr, val)
	}
	return sortedArr
}

func main() {
	nums := []int{12, 31, 411, 21, 1}
	sortedArr := sortFunc(nums)
	fmt.Println(sortedArr)
}
