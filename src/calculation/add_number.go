package calculation

// n数之和
func AddNumber(n int) (res int) {
	
	for i := 1; i <= n; i++ {
		res += i
	}
	
	return res
}

// 公式版
func AddNumber2(n int) int {
	return (1 + n) * n / 2
}