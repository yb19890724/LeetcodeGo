package sort

import (
	`github.com/hyb/LeetCode/src/sort`
	`github.com/stretchr/testify/assert`
	`testing`
)

// @test bubble sort
func TestBubbleSort(t *testing.T) {
	
	nilValue, emptyArray := sort.BubbleSort(nil), sort.BubbleSort([]int{})
	
	assert.Equal(t, nilValue, emptyArray, " data is nil value or is emptyArray sort error ")
	
	result := sort.BubbleSort([]int{1, 34, 6, 7, 10})
	
	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data bubble sort error. ")
}
