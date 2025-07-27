package main_test

import (
	"testing"

	main "github.com/matheuscaet/go-algorithms"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	sortedArr := main.BubbleSort(arr)
	if !equal(sortedArr, expected) {
		t.Errorf("BubbleSort failed: expected %v, got %v", expected, sortedArr)
	}
}

// equal compares two slices of int for equality.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestKangaroo(t *testing.T) {
	expected := "YES"
	result := main.Kangaroo(14, 4, 98, 2)
	if result != expected {
		t.Errorf("Kangaroo failed: expected %v, got %v", expected, result)
	}
}
