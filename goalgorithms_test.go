package main_test

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	sortedArr := main.BubbleSort(arr)
	if !equal(sortedArr, expected) {
		t.Errorf("BubbleSort failed: expected %v, got %v", expected, sortedArr)
	}
}
