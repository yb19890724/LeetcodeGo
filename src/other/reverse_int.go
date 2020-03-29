package other

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31


// 1.取余的方式计算个位剩余数
// 2.每次循环增加位数
// 3.计算完当前为向左移动一位，用于下次计算
func ReverseInt(x int) int {
	
	var num, newNum int
	
	for x != 0 { // 直到x等于0，跳出循环
		
		a := x % 10
		
		newNum = (num*10) + a
		
		num = newNum
		x = x / 10
		
		if num > MaxInt32 || num < MinInt32 {
			return 0
		}
	}
	
	return num
}
