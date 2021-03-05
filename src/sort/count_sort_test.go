package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test count sort
func TestCountingSort(t *testing.T) {

	result := CountSort([]int{2, 3, 0, 2, 3, 0, 3, 5})

	assert.Equal(t, result, []int{0, 0, 2, 2, 3, 3, 3, 5}, " data count sort error")
}

// @test count sort1
func TestCountSort1(t *testing.T) {
	
	result := CountSort([]int{2, 3, 0, 2, 3, 0, 3, 5})
	
	assert.Equal(t, result, []int{0, 0, 2, 2, 3, 3, 3, 5}, " data count sort error")
}

// @test count sort2
func TestCountSort2(t *testing.T) {
	
	result := CountSort2([]int{2, 3, 0, 2, 3, 0, 3, 5})
	
	assert.Equal(t, result, []int{0, 0, 2, 2, 3, 3, 3, 5}, " data count sort error")
}
