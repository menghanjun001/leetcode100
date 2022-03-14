package leetcode

//https://leetcode-cn.com/problems/minimum-path-sum/
//状态转移方程：res[i][j]=min(res[i-1][j]+res[i][j], res[i][j-1]+res[i][j])
func minPathSum(grid [][]int) int {
	//0.初始化

	//1.动态规划
	for i := range grid {
		for j := range grid[i] {
			if i == j && i == 0 { //起点
				continue
			} else if i == 0 { //上边
				grid[i][j] += grid[i][j-1]
			} else if j == 0 { //左边
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}
