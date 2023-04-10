package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// maxProfit 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	minPrice := 1<<63 - 1
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}
	return maxProfit
}

// repeatedSubstringPattern 338. 比特位计数
func countBits(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	// "{[()]}"
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func firstUniqChar(s string) byte {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if p, ok := m[target-nums[i]]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}

func isValid1(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 2 1 -1
func pivotIndex(nums []int) int {
	length := len(nums)
	pre := make([]int, length)
	suf := make([]int, length)
	pre[0] = nums[0]
	for i := 1; i < length; i++ {
		pre[i] = pre[i-1] + nums[i]
	}
	suf[length-1] = nums[length-1]
	for i := length - 1; i >= 1; i-- {
		suf[i-1] = suf[i] + nums[i-1]
	}
	for i := 0; i < length; i++ {
		if pre[i] == suf[i] {
			return i
		}
	}

	return -1
}

func main() {
	res := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	fmt.Println(res)
}
