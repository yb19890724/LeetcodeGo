package other

// 斐波那数列
// 递归
func Fibonacci(n int) int {
	
	if n <= 1 {
		return n
	}
	
	return Fibonacci(n-1) + Fibonacci(n-2)
}
