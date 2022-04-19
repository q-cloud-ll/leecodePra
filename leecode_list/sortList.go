package leecode_list

// 自底向上归并排序

// merge 合并两个链表
func Merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	// 遍历两个链表
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

//func sortList(head *ListNode) *ListNode {
//
//}
