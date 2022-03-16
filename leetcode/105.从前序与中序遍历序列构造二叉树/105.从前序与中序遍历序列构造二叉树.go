package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//preorder pop出来的总是根节点，根据根节点对中序遍历做分割，区分左右子树进行递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(&preorder, inorder)
}

func build(preorder *[]int, inorder []int) *TreeNode {
	//0.边界条件
	var (
		e    int
		idx  int
		root *TreeNode
	)
	if len(*preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//1.递归
	*preorder, e = pop(*preorder) //todo 递归中的preorder要改
	idx = findIndex(inorder, e)
	root = &TreeNode{
		Val: e,
	}
	root.Left = build(preorder, inorder[:idx])
	root.Right = build(preorder, inorder[idx+1:])
	return root
}

func findIndex(slice []int, e int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == e {
			return i
		}
	}
	return -1
}

func pop(slice []int) ([]int, int) {
	e := slice[0]
	tempSlice := make([]int, 0)
	tempSlice = append(tempSlice, slice[1:]...)
	return slice[1:], e
}
