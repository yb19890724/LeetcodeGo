package other

import (
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/other"
	"testing"
)

func TestFibonacci(t *testing.T) {
	
	data := [10]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	
	for i := 1; i <= 10; i++ {
		
		assert.Equal(t, other.Fibonacci(i), data[i-1], " fibonacci error")
		
	}
}
