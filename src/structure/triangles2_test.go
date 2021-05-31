package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test triangles
func TestTriangles2(t *testing.T) {
	
	res := []int{
		1, 3, 3, 1,
	}
	
	assert.Equal(t, Triangles2(3), res, " triangles struct failure ")
}
