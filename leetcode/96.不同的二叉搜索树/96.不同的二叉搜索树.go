package leetcode

//https://leetcode-cn.com/problems/unique-binary-search-trees/
//当前节点的种类=左子树种类*右子树种类
//f(i)=f(i-1)*f(n-i)
//递归写法
//func numTrees(n int) int {
//	var res = 0
//	if n == 0 || n == 1 {
//		return 1
//	}
//	for i := 1; i < n+1; i++ {
//		res += numTrees(i-1) * numTrees(n-i)
//	}
//	return res
//}
//dp写法
//f(i)=f(i-1)*f(n-i)
func numTrees(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ { //遍历dp
		for j := 1; j < i+1; j++ { //对于1组成的tree 对于1，2组成的tree 对于1,2,3...组成的tree
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
