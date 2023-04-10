package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxLength(nums []int) int {
	count, sum := 0, 0
	hm := map[int]int{}
	hm[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		sum += nums[i]
		if _, ok := hm[sum]; ok {
			count = max(count, i-hm[sum])
		} else {
			hm[sum] = i
		}
	}

	return count
}

// 前缀和+哈希表优化
func subarraySum(nums []int, k int) int {
	// [j...i] pre[i] - pre[j-1] == k
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre]++

	}

	return count
}

func findMaxLength1(nums []int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i] = -1
		}
		pre += nums[i]
		if _, ok := m[pre]; ok {
			count = max(count, i-m[pre])
		} else {
			m[pre] = i
		}

	}
	return count
}

func main() {
	res := findMaxLength1([]int{0, 0, 1, 0, 1})
	fmt.Println(res)
}
