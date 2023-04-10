package main

import (
	"fmt"
	"math"
	"sort"
)

//// 快速排序入口函数
//func quickSort(nums []int, p int, r int) {
//	// 递归终止条件
//	if p >= r {
//		return
//	}
//	// 获取分区位置
//	q := partition(nums, p, r)
//	// 递归分区（排序是在定位 pivot 的过程中实现的）
//	quickSort(nums, p, q-1)
//	quickSort(nums, q+1, r)
//}
//
//// 定位 pivot
//func partition(nums []int, p int, r int) int {
//	// 以当前数据序列最后一个元素作为初始 pivot
//	pivot := nums[r]
//	// 初始化 i、j 下标
//	i := p
//	// 后移 j 下标的遍历过程
//	for j := p; j < r; j++ {
//		// 将比 pivot 小的数丢到 [p...i-1] 中，剩下的 [i...j] 区间都是比 pivot 大的
//		if nums[j] < pivot {
//			// 互换 i、j 下标对应数据
//			nums[i], nums[j] = nums[j], nums[i]
//			// 将 i 下标后移一位
//			i++
//		}
//	}
//
//	// 最后将 pivot 与 i 下标对应数据值互换
//	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
//	nums[i], nums[r] = nums[r], nums[i]
//	// 返回 i 作为 pivot 分区位置
//	return i
//}

// 有一个整型的无序数组，如何快速找出第k大的数？
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func topKFrequent(nums []int, k int) []int {
	var res []int
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
			res = append(res, nums[i])
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return mp[res[i]] > mp[res[j]]
	})
	return res[:k]
}

func main() {
	var t, a, k int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &k)
		n := int(math.Pow(10, float64(k))) - a
		numerator, denominator := a, n
		numerator, denominator = simplifyFraction(numerator, denominator)
		fmt.Printf("%d/%d\n", numerator, denominator)
	}

}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 将分数化为最简分数
func simplifyFraction(numerator, denominator int) (int, int) {
	d := gcd(numerator, denominator)
	return numerator / d, denominator / d
}
