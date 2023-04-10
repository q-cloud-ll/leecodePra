package main

func MaxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// canJump 55. 跳跃游戏
func canJump(nums []int) bool {
	var p int
	for i := range nums {
		if i > p {
			return false
		}
		p = max(p, i+nums[i])
	}
	return true
}

// jump 跳跃游戏 II [2,3,1,2,4,2,3]
func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	step := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

// uniquePaths 62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}

// generate 118. 杨辉三角
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

// minPathSum 64. 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	// 先把第一行和第一列初始值相加好，后续dp
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

// climbStairs 70. 爬楼梯
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func getRow(rowIndex int) []int {
	dp := make([][]int, rowIndex+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp[rowIndex]
}

//func trap(height []int) int {
//	sum := 0
//	max_left := make([]int, len(height))
//	max_right := make([]int, len(height))
//
//	for i := 1; i < len(height)-1; i++ {
//		max_left[i] = max(max_left[i-1], height[i-1])
//	}
//}

func coinChange(coins []int, amount int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 1 << 32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func lengthOfLIS(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, len(nums))
	maxans := 1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxans = max(maxans, dp[i])
	}

	return maxans
}

func maxSubArray(nums []int) int {
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] += max(nums[i-1], 0)
		res = max(res, nums[i])
	}
	return res
}

//func longestCommonSubsequence(text1, text2 string) int {
//	m, n := len(text1), len(text2)
//	dp := make([][]int, m+1)
//	for i := 0; i <= m; i++ {
//		dp[i] = make([]int, n+1)
//	}
//	for i := 1; i < m+1; i++ {
//		for j := 1; j < n+1; j++ {
//			if text1[i-1] == text2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + 1
//			} else {
//				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//			}
//		}
//	}
//
//	return dp[m][n]
//}
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = append(dp[i], n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
