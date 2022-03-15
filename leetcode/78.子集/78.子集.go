package leetcode

//https://leetcode-cn.com/problems/subsets/
//dfs 剪枝+存入子集
func subsets(nums []int) [][]int {
	//0.初始化
	var (
		res    = make([][]int, 0)
		subset = make([]int, 0)
	)

	//1.dfs
	dfs(nums, &res, 0, subset)
	return res
}

func dfs(nums []int, res *[][]int, begin int, subset []int) {
	//0.加入子集
	*res = append(*res, subset)
	//1.结束递归
	if begin == len(nums) {
		return
	}

	//2.从begin遍历子节点，加入重置状态
	for i := begin; i < len(nums); i++ {
		subset = append(subset, nums[i])
		dfs(nums, res, i+1, subset)
		subset = pop(subset)
	}
}

func pop(src []int) []int {
	tgt := make([]int, 0)
	tgt = append(tgt, src[:len(src)-1]...)
	return tgt
}
