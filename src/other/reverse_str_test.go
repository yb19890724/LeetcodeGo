package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


// @test reverse str
func TestReverseStr(t *testing.T) {
	
	assert.Equal(t, ReverseStr(""), ""," reverse error")
	
	assert.Equal(t, ReverseStr("abc"), "cba"," reverse error")
	
	assert.Equal(t, ReverseStr("123"), "321"," reverse error")
	
	assert.Equal(t, ReverseStr2(""), ""," reverse error")
	
	assert.Equal(t, ReverseStr2("abc"), "cba"," reverse error")
	
	assert.Equal(t, ReverseStr2("123"), "321"," reverse error")
	
}
