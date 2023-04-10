package main

// binarySearch 二分查找
func binarySearch(nums []int, num int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// binarySearchV1 查找第一个值等于给定值的元素
func binarySearchV1(nums []int, num int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > num {
			right = mid - 1
		} else if nums[mid] < num {
			left = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != num {
				return mid
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// binarySearchV2 二分查找变形版本2：查找最后一个值等于给定值的元素
func binarySearchV2(nums []int, num int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > num {
			right = mid - 1
		} else if nums[mid] < num {
			left = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != num {
				return mid
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
