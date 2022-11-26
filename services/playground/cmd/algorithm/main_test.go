package algorithm

import "testing"

func TestLinearSearch(t *testing.T) {
	nums := []int{1, 4, 3, 7, 4, 6, 10}

	if LinearSearch(nums, 7) != 3 {
		t.Errorf("LinearSearch() = %v, want %v", LinearSearch(nums, 7), 3)
	}

	if LinearSearch(nums, 100) != -1 {
		t.Errorf("LinearSearch() = %v, want %v", LinearSearch(nums, 100), -1)
	}
}
