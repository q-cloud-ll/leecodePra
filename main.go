package main

import (
	"LeecodePra/leecode_list"
	"fmt"
)

func main() {
	res := minSwaps([]int{2, 4, 1, 3, 5}, 3)
	fmt.Println(res)
}
func minSwaps(arr []int, k int) int {
	n := len(arr)
	kCount := 0
	for i := 0; i < n; i++ {
		if arr[i] == k {
			kCount++
		}
	}
	if kCount == 0 {
		return 0
	}
	kPos := n - kCount
	swaps := 0
	for i := 0; i < n && kPos < n; i++ {
		if arr[i] != k {
			if i < kPos {
				kPos++
			} else {
				arr[kPos], arr[i] = arr[i], arr[kPos]
				swaps++
				kPos++
			}
		}
	}
	return swaps
}
func createList(list []int) *leecode_list.ListNode {
	root := &leecode_list.ListNode{Val: list[0]}
	tail := root
	for i := 1; i < len(list); i++ {
		tail.Next = &leecode_list.ListNode{Val: list[i]}
		tail = tail.Next
	}
	return root
}

func PrintList(list *leecode_list.ListNode) {
	p := list
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
