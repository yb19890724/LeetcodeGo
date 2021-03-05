package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test shell sort
func TestShellSort(t *testing.T) {
	
	result := ShellSort([]int{11,10,9,8,7,6,5,4,3,2,1})
	
	assert.Equal(t, result, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, " data insert sort error. ")
}