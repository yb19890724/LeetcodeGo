package structure

import "fmt"

// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
// 请你实现这个将字符串进行指定行数变换的函数：

//  m 根据行数分配大小 记录每一行放的内容
func ZConvert(s string, numRows int) string {
	
	// 不满足，提前返回
	if len(s) <= 2 || numRows == 1 {
		return s
	}
	// 保存最终结果
	var res string
	// 保存每次的结果
	var m = make([]string, numRows)
	// 初始位置
	curPos := 0
	// 用来标记是否该转向了
	shouldTurn := -1
	// 遍历每个元素
	for _, val := range s {
		// 添加进m里面，位置为curPos
		m[curPos] += string(val)
		// 如果走到头或者尾，就该转向了
		// 因为就在numRows的长度里面左右震荡
		if curPos == 0 || curPos == numRows-1 {
			// 转向
			shouldTurn = -shouldTurn
		}
		// curPos 用来标记m里面我们该去哪个位置
		curPos += shouldTurn
		
		fmt.Println(curPos)
	}
	// 这时m里面已经保存了数据了，我们只需要转换一下输出格式
	for _, val := range m {
		res += val
	}
	// 最后的输出
	return res
	
}