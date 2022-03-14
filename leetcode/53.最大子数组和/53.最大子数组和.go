package leetcode

//https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/
//动态规划
func maxSubArray(nums []int) int {
	//0.初始化 边界条件
	var (
		max = nums[0]
	)

	//1.遍历一次，根据上个数的状态判断当前数的状态
	for i := 1; i < len(nums); i++ {
		//判断加不加
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		//判断大小
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
