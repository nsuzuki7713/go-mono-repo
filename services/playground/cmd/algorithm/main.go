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

// 二分探索
func BinarySearch(nums []int, value int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[mid] == value {
			return mid
		}
		if nums[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// バブルソート
// 隣り合う要素の大小を比較しながら整列させるソートアルゴリズム。
func BubbleSort(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		for j := 0; j < lenNums-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
