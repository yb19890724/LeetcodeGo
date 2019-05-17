package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test selection sort
func TestSelectionSort(t *testing.T)  {

	result := SelectionSort([]int{1, 34, 6, 7, 10})

	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data selection sort error. ")
}

