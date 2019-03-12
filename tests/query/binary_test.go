package query

import (
	`github.com/hyb/LeetCode/src/query`
	`github.com/stretchr/testify/assert`
	`testing`
)

// @test binary search
func TestBinarySearch(t *testing.T) {
	
	exists := query.BinarySort([]int{1, 6, 7, 10, 34}, 10)
	
	assert.Equal(t, exists, 3, "Can't get the value that exists")
	
	notExists := query.BinarySort([]int{1, 6, 7, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, "Get nonexistent values")
}

// @test first binary search
func TestFirstBinarySearch(t *testing.T) {
	
	location := query.FirstBinarySearch([]int{1, 6, 10, 10, 34}, 10)
	
	assert.Equal(t, location, 2, "Get location Wrong")
	
	notExists := query.FirstBinarySearch([]int{1, 6, 10, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, "Get nonexistent values")
}

// @test last binary search
func TestLastBinarySearch(t *testing.T) {
	
	location := query.LastBinarySearch([]int{1, 6, 10, 10, 34}, 10)

	assert.Equal(t, location, 3, "Get location Wrong")
	
	notExists := query.LastBinarySearch([]int{1, 6, 10, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, "Get nonexistent values")
}