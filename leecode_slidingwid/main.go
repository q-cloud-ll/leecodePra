package main

import "fmt"

func lidingWindows(nums []int, target int) int {
	var left, right, ans int
	// 窗口初始化大小为 1，可以认为 right 已经加入窗口
	for right < len(nums) {
		// 处理 right 加入窗口后的逻辑
		//	...
		for { //判断是否需要收缩 {
			// 判断当前窗口大小是否满足题意
			if right-left > ans {
				ans = right - left + 1
			}
			// 处理 left 移除窗口后的逻辑
			left++ // shrink windows
		}
		right++ // scale windows
	}
	return ans
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	mp := map[byte]int{}
	m := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok {
			left = max(left, mp[s[i]]+1)
		}
		mp[s[i]] = i
		m = max(m, i-left+1)
	}
	return m
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	left, right, mult, ans := 0, 0, 1, 0
	for right < len(nums) {
		mult *= nums[right]
		for mult >= k {
			mult /= nums[left]
			left++
		}
		ans += right - left + 1
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}

	var cnt1, cnt2 [26]int
	for k, v := range s1 {
		cnt1[v-'a']++
		cnt2[s2[k]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func findAnagrams(s1, s2 string) []int {
	var ans []int
	n, m := len(s1), len(s2)
	if n > m {
		return []int{}
	}

	var cnt1, cnt2 [26]int
	for k, v := range s1 {
		cnt1[v-'a']++
		cnt2[s2[k]-'a']++
	}
	if cnt1 == cnt2 {
		ans = append(ans, 0)
		return ans
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func main() {
	res := lengthOfLongestSubstring("pwwkew")
	fmt.Println(res)
}
