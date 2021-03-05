package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// @test find median sorted arrays
func TestFindMedianSortedArrays(t *testing.T) {
	
	assert.Equal(t, FindMedianSortedArrays([]int{1,3},[]int{2}), 2.0, " max check for errors ")
	
	assert.Equal(t, FindMedianSortedArrays([]int{1,1},[]int{1,2}), 1.0, " max check for errors ")
}

