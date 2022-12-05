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

// 選択ソート
// 配列から最小値を探し、配列の先頭要素と入れ替えていくことで並べ替える。
func SelectionSort(nums []int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		minIdx := i
		for j := i + 1; j < lenNums; j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}

		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}

// 挿入ソート
func InsertionSort(nums []int) []int {
	lenNums := len(nums)

	for i := 1; i < lenNums; i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[j+1] = temp
	}

	return nums
}
