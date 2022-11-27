package algorithm

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	nums := []int{1, 4, 3, 7, 4, 6, 10}

	if LinearSearch(nums, 7) != 3 {
		t.Errorf("LinearSearch() = %v, want %v", LinearSearch(nums, 7), 3)
	}

	if LinearSearch(nums, 100) != -1 {
		t.Errorf("LinearSearch() = %v, want %v", LinearSearch(nums, 100), -1)
	}
}

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 7, 12, 24, 36, 37, 40}

	if BinarySearch(nums, 7) != 2 {
		t.Errorf("BinarySearch() = %v, want %v", BinarySearch(nums, 7), 2)
	}

	if BinarySearch(nums, 13) != -1 {
		t.Errorf("BinarySearch() = %v, want %v", BinarySearch(nums, 100), -1)
	}
}

func TestBubbleSort(t *testing.T) {
	nums := []int{1, 2, 4, 3, 7, 6, 5, 8, 9}

	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if r := BubbleSort(nums); !reflect.DeepEqual(r, expect) {
		t.Errorf("error")
	}
}
