package leetcode

func maxArea(height []int) int {
	//0.参数校验
	if len(height) == 0 {
		return 0
	}
	var (
		left, right = 0, len(height) - 1
		max         = 0
	)
	//1.循环，直至要重合
	for left<right {
		//3.记录当前水量，容积变大则更新标记位
		if capacity(left, right, height) > max {
			max = capacity(left, right, height)
		}
		//2.移动高度小的指针
		if height[left]<height[right] {
			left++
		}else{
			right--
		}
	}
	return max
}

func capacity(i int, j int, height []int) int {
	return min(height[i], height[j]) * (j - i)
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}
