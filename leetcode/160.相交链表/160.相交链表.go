package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
//https://leetcode-cn.com/problems/intersection-of-two-linked-lists/submissions/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var (
		pa, pb     = headA, headB
		lenA, lenB = 0, 0
	)

	for pa != nil || pb != nil {
		if pa != nil {
			lenA++
			pa = pa.Next
		}
		if pb != nil {
			lenB++
			pb = pb.Next
		}
	}
	pa, pb = headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			pa=pa.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			pb=pb.Next
		}
	}

	for pa != nil && pb != nil {
		if pa==pb {
			return pa
		}
		pa=pa.Next
		pb=pb.Next
	}
	return nil
}
