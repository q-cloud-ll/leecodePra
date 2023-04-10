package leecode_list

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return current
}

func reversePrint(head *ListNode) []int {
	ans := 0
	temp := head
	for temp != nil {
		ans++
		temp = temp.Next
	}
	nums := make([]int, ans)
	for head != nil {
		nums[ans-1] = head.Val
		head = head.Next
		ans--
	}
	return nums
}

func reverseListDemo(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//
//func copyRandomList(head *Node) *Node {
//	if head != nil {
//		return nil
//	}
//	for node := head; node != nil; node = node.Next.Next {
//		node.Next = &Node{Val: head.Val, Next: node.Next}
//	}
//}
