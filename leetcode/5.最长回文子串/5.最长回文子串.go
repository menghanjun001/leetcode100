package leetcode

//todo 有错误
func longestPalindrome(s string) string {
	var (
		reverseS    = ""
		res         [][]int
		maxLen      = 0
		maxEnd      = 0
		left, right = 0, 0
	)
	//初始化二维数组
	res = make([][]int, len(s))
	for i, _ := range res {
		res[i] = make([]int, len(s))
	}

	//1.参数校验 空字符串和长度为1的字符串就是本身
	if s == "" || len(s) == 1 {
		return s
	}
	//2.reverse 原字符串
	reverseS = reverse(s)
	//3.遍历源字符串和倒叙字符串的二维数组
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(reverseS); j++ {
			//4.如果相等则为1或加1，记录该位置,推测左右边界
			if s[i] == reverseS[j] {
				if i == 0 || j == 0 {
					res[i][j] = 1
				} else {
					res[i][j] = res[i-1][j-1] + 1
				}
				if res[i][j] > maxLen {
					maxLen = res[i][j]
					maxEnd = i
				}
			}
		}
	}
	//5.根据最终位置和记录最大长度确认左右边界
	right = maxEnd
	left = maxEnd - maxLen + 1 //从右往左数
	return s[left:right+1]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func reverse(s string) string {
	var (
		slice = []rune(s)
	)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}
