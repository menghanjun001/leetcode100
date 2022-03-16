package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	//0.边界条件
	if root==nil {
		return 0
	}
	//1.递归
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(i int, j int) int {
	if i>j {
		return i
	}
	return j
}
