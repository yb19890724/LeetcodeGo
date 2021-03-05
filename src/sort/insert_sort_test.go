package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test insert sort
func TestInsertSort(t *testing.T) {
	
	result := InsertSort([]int{4, 2, 9, 1})
	
	assert.Equal(t, result, []int{1, 2, 4, 9}, " data insert sort error. ")
}