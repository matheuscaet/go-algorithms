package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Sorting an array using bubble sort...")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	sortedArr := bubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)
	fmt.Println("Sorting completed!")
	fmt.Println("Goodbye!")

	fmt.Println("Reversing the sorted array...")
	reversedArr := reverse(sortedArr)
	fmt.Println("Reversed array:", reversedArr)
	fmt.Println("Reversing completed!")
	fmt.Println("Goodbye again!")
}

// Bubble sort implementation in Go
// This function sorts an array of integers using the bubble sort algorithm.
// It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func reverse(arr []int) []int {
	emptyArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		emptyArr[i] = arr[len(arr)-1-i]
	}
	return emptyArr
}
