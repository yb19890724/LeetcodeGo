package other

import (
	"strings"
)

const (
	INT_MAX = 1<<31 - 1
	INT_MIN = -1 << 31
)


// 假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。
// 如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231)

func MyAtoi(str string) int {
	
	pos := 1
	res := 0
	
	str = strings.TrimSpace(str) // 过滤空格
	
	if len(str) == 0 {
		return res
	}
	
	i, pos := isPlusMinus(str)
	

	for ; i < len(str); i++ {
		
		if oveflow, res := isOverflow(res * pos); oveflow {
			return res
		}
		
		// 如果第一个出现的不是数字直接返回
		if str[i] < '0' || string(str[i]) > "9" {
			return res * pos
		}
		
		// res *10 进位操作
		// '0' ascii码 十进制的值是48
		res = res*10 + int(str[i] - '0')
	}
	
	_,res = isOverflow(res * pos)
	
	return res
}

// 判断是否有正负符号，定位索引位置
func isPlusMinus(str string) (int, int) {
	i, pos := 0, 1
	
	// 保存翻转后的正负值
	if str[i] == '-' {
		pos = -1
		i++
	}
	
	// 如果带符号从后一位 翻转
	if str[i] == '+' {
		i++
	}
	
	return i, pos
}

// 判断是否溢出  INT_MAX (2^31 - 1) 或 INT_MIN (-2^31)
func isOverflow(res int) (bool, int) {
	
	if res >= INT_MAX {
		return true, INT_MAX
	}
	
	if res <= INT_MIN {
		return true, INT_MIN
	}
	return false, res
}



