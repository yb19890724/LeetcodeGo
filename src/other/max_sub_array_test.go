package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test test max sub array
func TestMaxSubArray(t *testing.T)  {

	assert.Equal(t, MaxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}), 6," max sub array sum error ")

	assert.Equal(t, DpMaxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}), 6," max sub array sum error  ")

}