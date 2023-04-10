package main

import (
	"fmt"
	"math"
	"sort"
)

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return 0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProduct(words []string) int {
	sum := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			b := diffStr(words[i], words[j])
			if b {
				sum = max(sum, len(words[i])*len(words[j]))
			}
		}
	}
	return sum
}

func diffStr(words1, words2 string) bool {
	s1 := []byte(words1)
	s2 := []byte(words2)
	for _, v := range s1 {
		for _, b := range s2 {
			if v == b {
				return false
			}
		}
	}
	return true
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i != j {
				if numbers[i]+numbers[j] == target {
					return []int{i, j}
				}
			}
		}
	}
	return nil
}

func twoSumPointer(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			return []int{l, r}
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return nil
}

func subarraySum(nums []int, k int) int {
	ans := 0
	b := make([]int, len(nums)+1)
	copy(b, nums)
	for i := 0; i < len(b); i++ {
		if i != len(b)-1 {
			if b[i]+b[i+1] == k {
				ans++
			}
		} else {
			return ans
		}
	}
	return ans
}

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}

	return
}

// 1 2 3       3 6 9
// 4 5 6       2 5 8
// 7 8 9       1 4 7
func rotate(matrix [][]int) [][]int {
	res := make([][]int, len(matrix))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(res))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res[i][j] = matrix[j][len(matrix)-i-1]
			//02
			//12
			//22
		}
	}

	return res
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			mid = right
		} else {
			mid = left + 1
		}
	}

	return left
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[:m], nums2)
	sort.Ints(nums1)
}

var mes chan string

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func main() {
	res := minSubArrayLen(4, []int{1, 4, 4})
	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
