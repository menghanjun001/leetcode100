package leetcode

import "testing"

func TestIsSymmetric(t *testing.T) {
	isSymmetric(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	})
}
