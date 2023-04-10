package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	l, r := len(matrix)-1, 0
	for l >= 0 && r < len(matrix[0]) {
		if matrix[l][r] > target {
			l--
		} else if matrix[l][r] < target {
			r++
		} else {
			return true
		}
	}
	return false
}

func main() {

}
