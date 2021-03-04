package calculation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test number og matches
func TestNumberOfMatches(t *testing.T) {
	assert.Equal(t, 1, numberOfMatches(2), " number og matches not equal ")
	assert.Equal(t, 0, numberOfMatches(1), " number og matches not equal ")

}

// @test number og matches1
func TestNumberOfMatches1(t *testing.T) {
	assert.Equal(t, 6, numberOfMatches1(7), " number og matches not equal ")
	assert.Equal(t, 7, numberOfMatches1(8), " number og matches not equal ")
}
