package leecode_sword_to_offer

// 二分法
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0 // 从左下角找
	for i >= 0 && j < len(matrix[0]) {
		// 如果目标值比这个值大，直接这一列都不要了，列往右
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] { // 如果目标值比这个值小，直接这一行不要了，行向上
			i--
		} else {
			return true
		}
	}
	return false
}
