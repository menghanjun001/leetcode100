package leetcode

import "testing"

func TestMaxPathSum(t *testing.T) {
	maxPathSum(&TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	})
}

func TestMaxPathSum1(t *testing.T) {
	maxPathSum(&TreeNode{
		Val:   -3,
	})
}