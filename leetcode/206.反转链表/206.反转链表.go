package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
//双指针做法,记录cur和pre
func reverseList(head *ListNode) *ListNode {
	//0.border and initial
	var (
		cur, pre *ListNode
	)

	//1.move two pointers
	pre = head
	for pre != nil {
		//temp
		temp := pre.Next
		//change direction
		pre.Next = cur
		//move pointers
		cur = pre
		pre = temp
	}
	return cur
}
