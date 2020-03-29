package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// @test reverse int
func TestReverseInt(t *testing.T) {
	
	assert.Equal(t, ReverseInt(123), 321," reverse error")
	
	assert.Equal(t, ReverseInt(0), 0," reverse error")
	
	assert.Equal(t, ReverseInt(2147483648), 0," reverse error")
	
	assert.Equal(t, ReverseInt(-2147483649), 0," reverse error")
	
}
