package leetcode

import "testing"

func TestQuickSort(t *testing.T) {
	slice := []int{1, 10, 11, 9, 14, 3, 2, 20, 19, 43, 57, 3, 2}
	quickSort(slice,0,len(slice)-1)
}
