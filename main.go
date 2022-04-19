package main

import (
	"LeecodePra/leecode_list"
	"fmt"
)

func main() {
	var list = make([]int, 5, 10)
	var list2 = make([]int, 5, 10)
	for i := 0; i < len(list); i++ {
		list[i] = i * 2
		list2[i] = i + 1
	}

	PrintList(leecode_list.Merge(createList(list), createList(list2)))
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
