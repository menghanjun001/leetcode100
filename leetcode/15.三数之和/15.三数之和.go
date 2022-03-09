package leetcode

import "sort"

//https://leetcode-cn.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var (
		res = [][]int{}
	)
	//0.检查
	if len(nums) < 3 {
		return res
	}
	//1.排序
	sort.Ints(nums)
	//2.标记位+右侧的左右指针，标记位循环
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//3.剪枝（排序后对后面大元素没必要判断）
		if nums[i] > 0 {
			break
		}
		left := i + 1
		right := len(nums) - 1
		//4.左右指针循环
		for left < right {
			//5.确定三数之和
			if nums[i]+nums[left]+nums[right] == 0 {
				tempRes := []int{nums[i],nums[left],nums[right]}
				res = append(res, tempRes)
				//todo 移动指针也要剪枝，并且有边界(边界在前，短路与)
				for left<right&&nums[left]==nums[left+1] {
					left++
				}
				for left<right&&nums[right]==nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
