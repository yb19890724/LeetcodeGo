package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test bubble sort
func TestBubbleSort(t *testing.T) {

	result := BubbleSort([]int{1, 34, 6, 7, 10})

	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data bubble sort error. ")
}
