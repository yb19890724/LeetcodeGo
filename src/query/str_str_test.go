package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 实现一个strstr 函数
// 给定一个haystack字符串和一个needle字符串，在haystack字符串中找出needle字符串出现的第一个位置(从0开始)。不存在返回-1

// @test
func TestStrStr(t  *testing.T)  {

	assert.Equal(t, StrStr("aa","a"), 0, " ")
	//assert.Equal(t, StrStr("baa","a"), 1, " ")

}
