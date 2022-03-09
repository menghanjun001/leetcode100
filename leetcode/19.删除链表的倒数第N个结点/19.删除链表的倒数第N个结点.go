package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
//https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//0.确定双指针
	var (
		left, right = head, head
	)
	left = head
	for n != 0 {
		right = right.Next
		n--
	}
	//todo 越界检查，如果是[1]，1的情况则说明该去除左节点
	if right==nil {
		return head.Next
	}
	//1.循环至右指针没有next
	for right.Next!=nil{
		left=left.Next
		right=right.Next
	}
	//2.左指针直接指向右指针
	left.Next=left.Next.Next
	return head
}
