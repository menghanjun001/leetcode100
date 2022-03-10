package leetcode

//https://leetcode-cn.com/problems/two-sum/
//双指针做法
//func twoSum(nums []int, target int) []int {
//	//0.边界
//	//1.左右指针
//	var (
//		res         = make([]int, 0)
//		left, right = 0, len(nums) - 1
//	)
//	//2.循环，移动指针找值
//	for left = 0; left < len(nums); left++ {
//		for right = left+1; right < len(nums); right++ {
//			if nums[left]+nums[right]==target {
//				return append(res, left,right)
//			}
//		}
//	}
//	return res
//}

//hash做法,需要查找就可以考虑hash
func twoSum(nums []int, target int) []int {
	//0.边界
	//1.hashmap k->v : value->index,通过匹配value可以找到下标
	var (
		m   = make(map[int]int)
		res = make([]int, 0)
	)
	//2.遍历，在hashmap里找target-e
	//- [2,7] 9
	//第一次存 2->0,即num[i]->i
	//第二次7->1 ，匹配上了
	for i := range nums {
		if targetI, exsit := m[target-nums[i]]; exsit {
			return append(res, targetI, i)
		}
		//3.没找到则放该元素e
		m[nums[i]] = i
	}
	return res
}
