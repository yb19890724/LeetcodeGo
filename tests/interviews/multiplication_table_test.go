package interviews

import (
	"github.com/stretchr/testify/assert"
	"github.com/yb19890724/leetcode-go/src/interviews"
	"testing"
)

// @test multiplication table
func TestMultiplicationTable(t *testing.T)  {

	result :=interviews.MultiplicationTable()

	assert.NotEmpty(t,result," result is nil ")

}