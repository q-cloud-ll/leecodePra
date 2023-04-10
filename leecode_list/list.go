package leecode_list

import (
	"sort"
)

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func ListLen(head *ListNode) int {
	dummyLen := head
	ans := 0
	for dummyLen != nil {
		ans++
		dummyLen = dummyLen.Next
	}
	return ans
}

func ReorderList(head *ListNode) *ListNode {
	var slice []int
	dummy := head
	for dummy != nil {
		slice = append(slice, dummy.Val)
		dummy = dummy.Next
	}
	cur := head
	l, r := 0, len(slice)-1
	for l <= r {
		if l != 0 { // 第一步先连右指针
			// 先连左指针
			cur.Next = &ListNode{Val: slice[l]}
			cur = cur.Next
		}
		if l == r {
			break
		}
		// 再连右指针
		cur.Next = &ListNode{Val: slice[r]}
		cur = cur.Next

		l++
		r--
	}
	return head
}

func MMergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var l1, l2 []int
	for node := list1; node != nil; node = node.Next {
		l1 = append(l1, node.Val)
	}
	for node := list2; node != nil; node = node.Next {
		l2 = append(l2, node.Val)
	}

	l1 = append(l1, l2...)
	sort.Slice(l1, func(i, j int) bool {
		if l1[i] < l1[j] {
			return true
		}
		return false
	})
	cur := &ListNode{}
	dummy := cur
	for i := 0; i < len(l1); i++ {
		dummy.Next = &ListNode{Val: l1[i]}
		dummy = dummy.Next
	}

	return cur.Next
}

//func hasCycle(head *ListNode) bool {
//	dummy := head
//	m := map[*ListNode]struct{}{}
//	for dummy != nil {
//		if _, ok := m[dummy]; ok {
//			return true
//		}
//		m[dummy] = struct{}{}
//		dummy = dummy.Next
//	}
//	return false
//}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	dummy := head
	slow, fast := dummy, dummy.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

func AAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := make([]int, 0), make([]int, 0)
	for node := l1; node != nil; node = node.Next {
		s1 = append(s1, node.Val)
	}
	for node := l2; node != nil; node = node.Next {
		s2 = append(s2, node.Val)
	}
	var s3 []int
	d := &ListNode{}
	cur := d

	if len(s1) > len(s2) {
		s3 = make([]int, 0)
		diff := len(s1) - len(s2)
		for i := 0; i < diff; i++ {
			s3 = append(s3, 0)
		}
		s3 = append(s3, s2...)
		for i := 0; i < len(s3); i++ {
			if i != len(s3)-1 {
				if s1[i+1]+s3[i+1] >= 10 {
					cur.Next = &ListNode{Val: s1[i] + s3[i] + 1}
				} else {
					cur.Next = &ListNode{Val: s1[i] + s3[i]}
				}
			} else {
				cur.Next = &ListNode{Val: s1[i] + s3[i]}
			}
			cur = cur.Next
		}
	} else {
		for i := 0; i < len(s1); i++ {
			if i != len(s1)-1 {
				if s1[i+1]+s2[i+1] >= 10 {
					cur.Next = &ListNode{Val: s1[i] + s2[i] + 1}
				} else {
					cur.Next = &ListNode{Val: s1[i] + s2[i]}
				}
			} else {
				cur.Next = &ListNode{Val: s1[i] + s2[i]}
			}
			cur = cur.Next
		}
	}
	//else {
	//	s3 = make([]int, len(s2))
	//	diff := len(s2) - len(s1)
	//	for i := 0; i < len(s2); i++ {
	//		if diff != 0 {
	//			s3[i] = 0
	//		}
	//		diff--
	//	}
	//	s3 = append(s3, s1...)
	//
	//}
	dummy := d
	for dummy != nil {
		dummy.Val = dummy.Val % 10
		dummy = dummy.Next
	}

	return d.Next
}

func GGetIntersectionNode(headA, headB *ListNode) *ListNode {
	mp := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		mp[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if mp[tmp] {
			return tmp
		}
	}

	return nil
}

func detectCycle(head *ListNode) *ListNode {
	mp := map[*ListNode]bool{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if mp[tmp] {
			return tmp
		}
		mp[tmp] = true
	}

	return nil
}
