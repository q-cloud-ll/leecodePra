package leecode_sword_to_offer

func minArray(numbers []int) int {
	min := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}
