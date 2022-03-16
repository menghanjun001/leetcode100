package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	//0.初始化
	var (
		slow, fast = head, head
	)
	//1.循环查找
	for {
		if fast==nil||fast.Next==nil {
			return false
		}
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			return true
		}
	}
}
