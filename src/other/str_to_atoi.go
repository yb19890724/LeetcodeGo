package other

import (
	"math"
	"strings"
)

// 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。
// 如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231)

func MyAtoi(str string) int {
	
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	res := 0
	for _, ch := range str {
		
		// 不是数字
		if ch < '0' || ch > '9' {
			break
		}
		
		//  res *10 进位操作
		// '0' ascii码 十进制的值是48
		res = res*10 + int(ch-'0')
		if res > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return res * sign
}

