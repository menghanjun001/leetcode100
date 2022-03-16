package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	//0.初始化 边界条件
	var (
		res = make([][]int, 0)
	)
	if root == nil {
		return res
	}
	//1.bfs
	bfs(root, &res)
	return res
}

func bfs(root *TreeNode, res *[][]int) {
	//0.初始化queue
	queue := make([]*TreeNode, 0)
	//1.入栈
	queue = append(queue, root)
	//2.循环不为空的栈
	for len(queue) != 0 {
		//todo 分层
		tempRes := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ { //queue此时改变了
			var node *TreeNode
			//3.pop
			queue, node = pop(queue)
			tempRes = append(tempRes, node.Val)
			//4.pop出来的元素加入queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		*res = append(*res, tempRes)
	}
	return
}

func pop(queue []*TreeNode) ([]*TreeNode, *TreeNode) {
	node := queue[0]
	temp := make([]*TreeNode, 0)
	temp = append(temp, queue[1:]...)
	return temp, node
}
