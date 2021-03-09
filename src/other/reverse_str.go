package other

// 解法1
// 循环字符串，按照字符串长度从右往左 移动字符串
func ReverseStr(s string) string {
	
	l := len(s)
	
	res := make([]byte,l)
	
	for i := 0; i < l; i++ {
		res[i] = s[l-i-1]
	}
	return string(res)
}

// 解法2
// 循环字符串，按照字符串长度从右往左
func ReverseStr2(s string) string {
	
	l := len(s)
	
	res := make([]byte,l)
	
	for i := l-1; i >= 0; i-- {
		res[l-i-1] = s[i]
	}
	return string(res)
}