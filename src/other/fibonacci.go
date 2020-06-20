package other

// 斐波那数列
// 递归
// 时间复杂度 O(2^n）
func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 非递归版本
// 时间复杂度 O(n)
func Fibonacci2(n int) int {
	
	if n <= 1  {
		return n
	}
	
	var (
		first int
 		second =1
	)
	
	for i := 0; i < n -1; i++ {
		s :=first + second
		first = second
		second = s
	}
	
	return second
}