package leecode_sword_to_offer

func findRepeatNumber(nums []int) int {
	var res int
	numMap := make(map[int]int)
	for k, _ := range nums {
		if _, ok := numMap[nums[k]]; ok {
			res = nums[k]
			break
		}
		numMap[nums[k]]++
	}
	return res
}
