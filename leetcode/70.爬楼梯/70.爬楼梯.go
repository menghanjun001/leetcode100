package test

//https://leetcode-cn.com/problems/climbing-stairs/
//动态规划 或 fibonacci
//res[n]=res[n-1]+res[n-2]
func climbStairs(n int) int {
	//0.边界条件和初始化
	var dp = make([]int, n+1) //从0到n都有值
	dp[0] = 1
	dp[1] = 1
	//1.dp
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
