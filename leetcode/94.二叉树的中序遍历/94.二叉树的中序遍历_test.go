package leetcode

import "testing"

func TestInorderTraversal(t *testing.T) {
	inorderTraversal(&TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	})
}
