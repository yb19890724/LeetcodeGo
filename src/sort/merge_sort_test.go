package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test insert sort
func TestMergeSort(t *testing.T) {
	
	data := []int{4, 2, 9, 1}
	result := MergeSort(data, 0, len(data))
	
	assert.Equal(t, result, []int{1, 2, 4, 9}, " data merge sort error. ")
}