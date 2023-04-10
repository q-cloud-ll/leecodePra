package main

import (
	"fmt"
)

// mergeSort 归并排序
//func mergeSort(nums []int) []int {
//	if len(nums) <= 1 {
//		return nums
//	}
//
//	p := len(nums) / 2
//	left := mergeSort(nums[0:p])
//	right := mergeSort(nums[p:])
//
//	return merge(left, right)
//}

//func merge(left []int, right []int) []int {
//	i, j := 0, 0
//	m, n := len(left), len(right)
//	var res []int
//	for {
//		if i >= m || j >= n {
//			break
//		}
//
//		if left[i] <= right[j] {
//
//		}
//	}
//}

func deferCall() int {
	defer func() {
		fmt.Println("123")
	}()
	defer func() {
		fmt.Println("456")
	}()
	panic("panic")
	return 1231
}
