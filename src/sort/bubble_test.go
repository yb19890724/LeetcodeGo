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


func TestBubbleSort2(t *testing.T) {
	
	result := BubbleSort2([]int{1, 34, 6, 7, 10})
	
	assert.Equal(t, result, []int{1, 6, 7, 10, 34}, " data bubble sort error. ")
}


// @test Sum Range
func BenchmarkBubbleSort(b *testing.B){
	for i:=0;i<b.N;i++{
		_=BubbleSort([]int{1, 34, 6, 7, 10})
	}
}

// @test Sum Range2
func BenchmarkBubbleSort2(b *testing.B){
	for i:=0;i<b.N;i++{
		_=BubbleSort2([]int{1, 34, 6, 7, 10})
	}
}