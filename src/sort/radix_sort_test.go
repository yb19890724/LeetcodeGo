package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test radix sort
func TestRadixSort(t *testing.T) {

	result := RadixSort([]int{2, 3, 0, 2, 3, 0, 3, 5})

	assert.Equal(t, result, []int{0, 0, 2, 2, 3, 3, 3, 5}, " data radix sort error")
}


func TestBucketSort(t *testing.T) {
	
	result := BucketSort([]int{ 61, 60, 62, 63, 64, 66, 65, 67, 68, 69, 70, 71})
	
	assert.Equal(t, result, []int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71}, " data radix sort error")
}
