package calculation

import (
	`github.com/hyb/LeetCode/src/calculation`
	`github.com/stretchr/testify/assert`
	`testing`
)

// 思路：判断数组中所有值得组合方式，是否满足条件

func TestTwoSum(t *testing.T) {
	
	result := calculation.TwoSum([]int{3, 2, 3}, 6)
	
	assert.Equal(t, result, []int{0, 2}, " value not equal ")
	
}
