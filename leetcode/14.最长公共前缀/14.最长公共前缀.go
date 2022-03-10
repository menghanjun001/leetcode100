package leetcode

//https://leetcode-cn.com/problems/longest-common-prefix/solution/zui-chang-gong-gong-qian-zhui-by-leetcode-solution/
func longestCommonPrefix(strs []string) string {
	//0.边界
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	//1.一层循环 取index
	for idx := 0; idx < len(strs[0]); idx++ {
		//2.二层循环 用第一条比其他条
		for j := 1; j < len(strs); j++ {
			//3.终止条件 两条长度相等 或 当前位元素不等, 便从第一条削减
			if idx == len(strs[j]) || strs[0][idx] != strs[j][idx] {
				return strs[j][:idx]
			}
		}
	}
	return strs[0]

}
