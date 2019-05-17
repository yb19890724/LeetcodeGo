package calculation

import (
	`github.com/stretchr/testify/assert`
	`testing`
)

// @test two sum value
func TestTwoSum(t *testing.T) {
	
	result :=TwoSum([]int{3, 2, 3}, 6)
	
	assert.Equal(t, result, []int{0,2}, "  two sum value not equal ")
	
}


