package leetcode

import "fmt"

var (
	used [][]bool
)
//todo 没写完
//https://leetcode-cn.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	//0.初始化
	used = make([][]bool, len(board))
	for i, _ := range used {
		used[i] = make([]bool, len(board[0]))
	}

	//1.dfs
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i int, j int, idx int) bool {
	fmt.Println(string(board[i][j]))
	//0.终止条件
	if board[i][j] != word[idx] {
		return false
	}
	if board[i][j] == word[idx] && used[i][j] == false && idx == len(word)-1 {
		return true
	}
	//1.进入标记
	used[i][j] = true
	//2.进入递归
	if canLeft(board, i, j) {
		dfs(board, word, i, j-1, idx+1)
	}
	if canRight(board, i, j) {
		dfs(board, word, i, j+1, idx+1)
	}
	if canUp(board, i, j) {
		dfs(board, word, i-1, j, idx+1)
	}
	if canDown(board, i, j) {
		dfs(board, word, i+1, j, idx+1)
	}
	//3.退出标记
	used[i][j] = false
	return false
}

func canDown(board [][]byte, i int, j int) bool {
	if i+1 < len(board) && i+1 >= 0 && used[i+1][j] == false {
		return true
	}
	return false
}

func canUp(board [][]byte, i int, j int) bool {
	if i-1 < len(board) && i-1 >= 0 && used[i-1][j] == false {
		return true
	}
	return false
}

func canRight(board [][]byte, i int, j int) bool {
	if j+1 < len(board[0]) && j+1 >= 0 && used[i][j+1] == false {
		return true
	}
	return false
}

func canLeft(board [][]byte, i int, j int) bool {
	if j-1 < len(board[0]) && j-1 >= 0 && used[i][j-1] == false {
		return true
	}
	return false
}
