package leecode_list

//func ReverseList(head *ListNode) *ListNode {
//	// 1->2->3->4->5->nil
//	// nil <- 1 <- 2 <- 3
//	// prev 1>nil
//	// curr 2
//	var prev *ListNode
//	cur := head
//	for cur != nil {
//		next := cur.Next
//		cur.Next = prev
//		prev = cur
//		cur = next
//	}
//	return prev
//}

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy

	var pre *ListNode
	var i = 0
	for i < m {
		pre = head
		head = head.Next
		i++
	}

	var j = i
	var next *ListNode
	var mid = head
	for head != nil && j <= n {
		temp := head.Next
		head.Next = next
		next = head
		head = temp
		j++
	}
	pre.Next = next
	mid.Next = head
	return dummy.Next
}

func DeleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	prev := dummy
	for head != nil {
		next := head.Next
		if head.Val == val {
			prev.Next = next
			break
		}
		prev = head
		head = next
	}
	return dummy.Next
}

func IsPalindrome(head *ListNode) bool {
	var vals []int
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

func HammingWeight(num uint32) int {
	var res uint32
	for num != 0 {
		res += num & 1
		num >>= 1
	}
	return int(res)
}

func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// LengthOfLongestSubstring 最长字串长度
func LengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		// i=0, (a)bcabcbb, n = 8
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tmp := dummy
	t := 0
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{Val: t % 10}
		tmp = tmp.Next
		t /= 10
	}
	return dummy.Next
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}

// GenerateParenthesis 括号生成
func GenerateParenthesis(n int) []string {
	res := new([]string)
	backtracking(n, n, "", res)
	return *res
}

func backtracking(left, right int, tmp string, res *[]string) {
	if right == 0 {
		*res = append(*res, tmp)
		return
	}
	if left > 0 {
		backtracking(left-1, right, tmp+"(", res)
	}
	if right > left {
		backtracking(left, right-1, tmp+")", res)
	}
}

// FindTheDifference 找不同
func FindTheDifference(s string, t string) byte {
	sum := 0
	for _, ch := range s {
		sum -= int(ch)
	}
	for _, ch := range t {
		sum += int(ch)
	}
	return byte(sum)
}

// SingleNumber 只出现一次的数字
func SingleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

// oddEvenList 328. 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// nodesBetweenCriticalPoints 2058. 找出临界点之间的最小和最大距离
//func nodesBetweenCriticalPoints(head *ListNode) []int {
//	if head == nil && head.Next == nil && head.Next.Next == nil {
//		return []int{-1, -1}
//	}
//	dummy := head
//	prev := dummy
//	next := dummy.Next
//	minDistance := -1
//	maxDistance := -1
//	for dummy != nil {
//		if prev.Val > {
//
//		}
//		dummy = dummy.Next
//	}
//}

func GetKthFromEnd(head *ListNode, k int) *ListNode {
	tmp := &ListNode{}
	n := 0
	node := head
	for node != nil {
		n++
		node = node.Next
	}
	tmp = head
	for ; n > k; n-- {
		tmp = tmp.Next
	}
	return tmp
}

// GetKthFromEnd2 快慢指针解法
func GetKthFromEnd2(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

//func ReverseList(head *ListNode) *ListNode {
//	// 1 > 2 > 3 > 4
//	var prev *ListNode
//
//}
