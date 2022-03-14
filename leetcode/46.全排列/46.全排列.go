package leetcode

//https://leetcode-cn.com/problems/permutations/
//todo 没写出来 草拟吗
func permute(nums []int) [][]int {
	//0.边界条件，初始化
	var (
		res  = make([][]int, 0)
		path = make([]int, 0)
	)

	//1.回溯
	dfs(nums, path, &res)
	return res
}

func dfs(nums []int, path []int, res *[][]int) {
	//0.边界条件
	if len(nums) == 0 {
		tempPath := path
		copy(tempPath, path)
		*res = append(*res, tempPath)
		return
	}
	for i := 0; i < len(nums); i++ {
		//1.递归
		tempNums,e := pop(nums, i)
		path = append(path, e)
		dfs(tempNums, path, res)
	}
}

//pop 返回弹出元素后的数组 与 弹出元素
//todo 坑：对于递归使用的一个slice，他们指向的数组可能是同一个，在递归中改变会一变而动全局
//https://www.codegrepper.com/code-examples/go/golang+pop+from+slice
func pop(nums []int, i int) (res []int, e int) {
	res = make([]int, 0) //slice的切片操作是移动指针，不是新建内存
	e = nums[i]
	res = append(res, nums[:i]...)
	res = append(res, nums[i+1:]...)
	return
}
