package algorithm

// 線形探索
func LinearSearch(nums []int, value int) int {
	for i, num := range nums {
		if num == value {
			return i
		}
	}

	return -1
}
