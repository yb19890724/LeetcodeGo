package sort

import (
	`github.com/hyb/LeetCode/src/sort`
	`github.com/stretchr/testify/assert`
	`testing`
)

// @test quick sort
func TestQuickSort(t *testing.T) {
	
	result := sort.QuickSort([]int{1, 34, 6, 7, 10})
	
	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " quick sort error")
}
