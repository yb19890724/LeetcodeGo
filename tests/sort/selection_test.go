package sort

import (
	"github.com/hyb/LeetCode-Go/src/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T)  {

	result := sort.QuickSort([]int{1, 34, 6, 7, 10})

	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data selection sort error. ")
}

