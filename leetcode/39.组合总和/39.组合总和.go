package leetcode

import "sort"

//https://leetcode-cn.com/problems/combination-sum/
//以target构建树,target-candidates[i]作为子节点，终止条件是叶子节点<0或=0，记录路径（candidates[i]）
func combinationSum(candidates []int, target int) [][]int {
	//0.初始化
	var (
		res  = make([][]int, 0)
		path = make([]int, 0)
	)
	//1.边界条件
	if len(candidates) == 0 {
		return res
	}
	//2.dfs,排序后设置起点可以剪枝
	sort.Ints(candidates)
	dfs(candidates,target, 0, path, &res)
	return res
}

func dfs(candidates []int, target int, begin int, path []int, res *[][]int) {
	//0.终止条件
	//-当前target<0，return
	//-当前target=0，记录该条path到res，return
	if target < 0 {
		path=path[:len(path)-1]
		return
	}
	if target == 0 {
		tempPath:=make([]int,len(path)) //todo 一次深拷贝，防止底层数组变化
		copy(tempPath,path)
		*res = append(*res, tempPath) //path需要深copy，数组内部会被改变
		return
	}
	//1.遍历子节点进行递归
	for i := begin; i < len(candidates); i++ {
		path = append(path, candidates[i])
		dfs(candidates, target-candidates[i], i, path, res)
		//2.状态重置,有重置才有回溯
		path=path[:len(path)-1]
	}
}
