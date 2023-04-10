package leecode_sword_to_offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	tmp := make([]int, 0)
	res := make([]int, 0)
	if head == nil {
		return nil
	}

	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	//for i := len(tmp) - 1; i >= 0; i++ {
	//	res = append(res, tmp[i])
	//}
	for len(tmp) > 0 {
		res = append(res, tmp[len(tmp)-1])
		tmp = tmp[:len(tmp)-1]
	}
	return res
}
