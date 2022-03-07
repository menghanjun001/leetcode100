package leetcode

func lengthOfLongestSubstring(s string) int {
	var (
		m                = make(map[byte]int) //存字符和次数
		left, right, ans = 0, -1, 0
	)
	for left = 0; left < len(s); left++ {
		//4.左指针移动一次就删掉已记录字符
		if left != 0 {
			delete(m, s[left-1])
		}

		for right+1 < len(s) && m[s[right+1]] == 0 { //1.如果右指针移动不越界 并 没存过的话
			//2.右指针一直移动更新，更新记录值
			right++
			m[s[right]]++
		}
		//3.更新长度
		ans = max(ans, right-left+1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x

	}
	return y
}
