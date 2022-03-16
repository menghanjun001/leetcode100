package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/solution/114-er-cha-shu-zhan-kai-wei-lian-biao-by-ming-zhi-/
func flatten(root *TreeNode) {
	//0.边界条件
	if root == nil {
		return
	}
	//1.递归
	//flattern左右子树
	flatten(root.Left)
	flatten(root.Right)
	//记录右子树
	temp := root.Right
	//左子树移动到右边
	root.Right = root.Left
	root.Left = nil
	//接上记录的右子树
	head := root
	for head.Right != nil {
		head = head.Right
	}
	head.Right=temp
}
