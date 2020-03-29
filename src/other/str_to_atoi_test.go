package other

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	
	"testing"
)

// @test str to atoi
func TestMyAtoi(t *testing.T) {
	
	assert.Equal(t, MyAtoi("42"), 42, " atoi error")
	assert.Equal(t, MyAtoi("   -42"), -42, " atoi error")
	assert.Equal(t, MyAtoi("words and 987"), 0, " atoi error")
	
	s :="0"
	
	fmt.Println(0x10) // 10进制
	fmt.Println(string('0')) // 10进制
	fmt.Println(int(s[0]-'0')) // 10进制
	
	
	
	fmt.Println(reflect.TypeOf(s[0]))
	fmt.Println(reflect.TypeOf(int(s[0])))
	fmt.Println(reflect.TypeOf('0'))
	fmt.Println(reflect.TypeOf(int(s[0])-'0'),s[0],'0',int(s[0])-'0')
	
	
	
}
