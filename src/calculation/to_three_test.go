package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// @test three sum value
func TestThreeSum(t *testing.T) {
	
	result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	
	assert.Equal(
		t,
		result,
		[][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
		"  two three value not equal ",
	)

}
