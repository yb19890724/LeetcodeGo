package sort

import (
	"github.com/yb19890724/leetcode-go/src/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test selection sort
func TestSelectionSort(t *testing.T)  {

	result := sort.QuickSort([]int{1, 34, 6, 7, 10})

	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data selection sort error. ")
}

