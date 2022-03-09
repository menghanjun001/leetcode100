package leetcode

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
	head:=ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	removeNthFromEnd(&head,2)
}

func TestRemoveNthFromEnd2(t *testing.T) {
	head:=ListNode{
		Val:  1,
		Next: nil,
	}
	removeNthFromEnd(&head,1)
}