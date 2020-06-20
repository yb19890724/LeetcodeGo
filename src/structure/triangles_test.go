package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test triangles
func TestTriangles(t *testing.T) {
	
	res := [][]int{
		{
			1,0,0,
		},
		{
			1,1,0,
		},
		{
			1,2,1,
		},
	}
	
	assert.Equal(t, Triangles(3), res, " triangles struct failure ")
}


