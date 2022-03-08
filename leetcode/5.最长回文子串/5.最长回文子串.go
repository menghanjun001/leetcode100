package leetcode


func longestPalindrome(s string) string {
	//边界条件
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var (
		left, right = 0, 0
		flag        [][]bool
	)
	//0.初始化+终止条件
	flag = make([][]bool, len(s))
	for i, _ := range flag {
		flag[i] = make([]bool, len(s))
		flag[i][i]=true
	}
	//1.双层循环，移动右左指针
	for j := 0; j < len(s); j++ {
		for i := 0; i < j; i++ {
			//2.转移条件：如果左指针的值=右指针的值
			if s[i] == s[j] {
				//- 左右间隔太小<2，即重合和相邻，则该子串是回文
				if j-i < 2 {
					flag[i][j] = true
				} else {
					//- 间隔大，左右内含子串的状态即当前子串的状态
					flag[i][j] = flag[i+1][j-1]
				}
				//3.如果是回文，并且长度更长，则更新左右标记位置
				if flag[i][j] == true && j-i+1 >= right-left+1 {
					left, right = i, j
				}
			}
		}
	}
	return s[left : right+1]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
