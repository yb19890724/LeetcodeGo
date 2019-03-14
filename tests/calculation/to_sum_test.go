package calculation

import (
	`github.com/hyb/LeetCode/src/calculation`
	`github.com/stretchr/testify/assert`
	`testing`
)


func TestTwoSum(t *testing.T) {
	
	result := calculation.TwoSum([]int{3, 2, 3}, 6)
	
	assert.Equal(t, result, []int{0, 2}, "  two sum value not equal ")
	
}
