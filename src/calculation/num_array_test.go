package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// @test sum range value
func TestSumRange(t *testing.T) {
	
	// [0 -2 -2 1 -4 -2 -3]
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	
	assert.Equal(t, 1 ,n.SumRange(0,2), " sum range value not equal ")
	assert.Equal(t, -1 ,n.SumRange(2,5), " sum range value not equal ")
	
	
	// [0 -2 -2 1 -4 -2 -3]
	n2 := Constructor2([]int{-2, 0, 3, -5, 2, -1})

	assert.Equal(t, 1 ,n2.SumRange2(0,2), " sum range value not equal ")
	assert.Equal(t, -1 ,n2.SumRange2(2,5), " sum range value not equal ")
	
}

// @test Sum Range
func BenchmarkSumRange(b *testing.B){
	n := Constructor([]int{-2, 0, 3, -5, 2, -1})
	for i:=0;i<b.N;i++{
		_=n.SumRange(0,2)
	}
}

// @test Sum Range2
func BenchmarkSumRange2(b *testing.B){
	n2 := Constructor2([]int{-2, 0, 3, -5, 2, -1})
	for i:=0;i<b.N;i++{
		_=n2.SumRange2(0,2)
	}
}