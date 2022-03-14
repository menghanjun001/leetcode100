package leetcode

//https://leetcode-cn.com/problems/unique-paths/
//动态规划,转移方程:res[i][j]=res[i-1][j]+res[i][j-1]

func uniquePaths(m int, n int) int {
	//0.初始化
	var (
		res = make([][]int, m)
	)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			if i == 0 || j == 0 {
				res[i][j] = 1
			}
		}
	}
	//1.动态规划
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}
