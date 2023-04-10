package leecode_sword_to_offer

func missingNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v] = 1
	}
	for i := 0; ; i++ {
		if numMap[i] == 0 {
			return i
		}
	}
}
