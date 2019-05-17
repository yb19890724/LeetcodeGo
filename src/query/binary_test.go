package query

import (
	`github.com/stretchr/testify/assert`
	`testing`
)

// @test binary search
func TestBinarySearch(t *testing.T) {
	
	exists := BinarySearch([]int{1, 6, 7, 10, 34}, 10)
	
	assert.Equal(t, exists, 3, " can't get the value that exists. ")
	
	notExists := BinarySearch([]int{1, 6, 7, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, " get nonexistent values. ")
}

// @test first binary search
func TestFirstBinarySearch(t *testing.T) {
	
	location := FirstBinarySearch([]int{1, 6, 10, 10, 34}, 10)
	
	assert.Equal(t, location, 2, " get location wrong. ")
	
	notExists := FirstBinarySearch([]int{1, 6, 10, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, " get nonexistent values. ")
}

// @test last binary search
func TestLastBinarySearch(t *testing.T) {
	
	location := LastBinarySearch([]int{1, 6, 10, 10, 34}, 10)

	assert.Equal(t, location, 3, " get location wrong. ")
	
	notExists := LastBinarySearch([]int{1, 6, 10, 10, 34}, 0)
	
	assert.Equal(t, notExists, -1, " get nonexistent values. ")
}