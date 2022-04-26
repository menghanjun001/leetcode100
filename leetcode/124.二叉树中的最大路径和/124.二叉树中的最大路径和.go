package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
//dfs，记录每个节点的最大贡献，遍历过程中更新最大路径和
func maxPathSum(root *TreeNode) int {
	//0.init
	var (
		maxSum = math.MinInt32 //todo 可能只有一个节点且为负数
		dfs    func(*TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		//0.border case
		if node == nil {
			return 0
		}
		var (
			leftGain, rightGain, newSum int
		)

		//1.递归计算左右节点的最大贡献值，负贡献记为0
		leftGain = max(dfs(node.Left), 0)
		rightGain = max(dfs(node.Right), 0)
		//2.获取当前节点的新路径和，更新最大路径和
		newSum = leftGain + node.Val + rightGain
		maxSum = max(newSum, maxSum)
		//3.返回当前节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	//1.dfs
	dfs(root)
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
