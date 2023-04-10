package leecode_list

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}

	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}

	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}

	return dummy.Next
}

func partition(head *ListNode) {

}
