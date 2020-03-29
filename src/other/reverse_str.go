package other

func ReverseStr(s string) string {
	
	l := len(s)
	
	res := make([]byte,l)
	
	for i := 0; i < l; i++ {
		res[i] = s[l-i-1]
	}
	return string(res)
}


func ReverseStr2(s string) string {
	
	l := len(s)
	
	res := make([]byte,l)
	
	for i := l-1; i >= 0; i-- {
		res[l-i-1] = s[i]
	}
	return string(res)
}