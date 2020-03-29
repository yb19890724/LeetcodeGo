package other

import (
	"github.com/stretchr/testify/assert"
	
	"testing"
)

// @test str to atoi
func TestMyAtoi(t *testing.T) {
	
	assert.Equal(t, MyAtoi("42"), 42, " atoi error")
	assert.Equal(t, MyAtoi("   -42"), -42, " atoi error")
	assert.Equal(t, MyAtoi("words and 987"), 0, " atoi error")
	
}
