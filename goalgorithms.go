package main

import (
	"fmt"
)

var dbData = []string{"id1", "id2", "id3", "id4"}

func main() {
	fmt.Println("Sorting an array using bubble sort...")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)
	sortedArr := BubbleSort(arr)
	fmt.Println("Sorted array:", sortedArr)

	fmt.Println("Reversing the sorted array...")
	reversedArr := Reverse(sortedArr)
	fmt.Println("Reversed array:", reversedArr)

	fmt.Println("Kangaroo problem...")
	fmt.Println(Kangaroo(0, 3, 4, 2))
}

// Bubble sort implementation in Go
// This function sorts an array of integers using the bubble sort algorithm.
// It repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order.
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func Reverse(arr []int) []int {
	emptyArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		emptyArr[i] = arr[len(arr)-1-i]
	}
	return emptyArr
}

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// If both kangaroos have the same velocity
	if v1 == v2 {
		// They will meet only if they start at the same position
		if x1 == x2 {
			return "YES"
		}
		return "NO"
	}

	// If kangaroo 1 is behind and moving slower, they will never meet
	if x1 < x2 && v1 <= v2 {
		return "NO"
	}

	// If kangaroo 1 is ahead and moving faster, they will never meet
	if x1 > x2 && v1 >= v2 {
		return "NO"
	}

	// Check if there's a valid integer solution for when they meet
	// x1 + v1*t = x2 + v2*t
	// (v1 - v2)*t = x2 - x1
	// t = (x2 - x1) / (v1 - v2)

	if (x2-x1)%(v1-v2) == 0 && (x2-x1)/(v1-v2) >= 0 {
		return "YES"
	}

	return "NO"
}
