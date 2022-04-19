package leecode_list

func RotateRight(head *ListNode, k int) *ListNode {
	// k == 0 头为空，头的下一个节点为空不用处理，返回 head 即可
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	// 看一下链表的长度是多少
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	// 看 k 是否是 n 的倍数，是倍数就不用处理
	add := n - k%n
	if add == n {
		return head
	}
	// 让iter的下一个节点连头，形成环状
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	res := iter.Next
	iter.Next = nil
	return res
}
