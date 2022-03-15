package leetcode

import "testing"

func TestExsit(t *testing.T) {
	src := make([][]byte, 3)
	src[0]=[]byte{'A','B','C','E'}
	src[1]=[]byte{'S','F','C','S'}
	src[2]=[]byte{'A','D','E','E'}
	exist(src, "ABCCED")
}
