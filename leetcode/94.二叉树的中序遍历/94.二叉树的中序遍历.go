package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	//0.边界条件
	if root == nil {
		return nil
	}
	//1.递归
	res := make([]int, 0)
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)
	return res
}
