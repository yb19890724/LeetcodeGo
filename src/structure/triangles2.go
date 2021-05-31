package structure

// 杨辉三角

func Triangles2(n int) []int {
	res := make([]int, n+1)
	res[0] = 1
	
	for i := 1; i < n+1; i++ {
		
		for j := i; j > 0; j-- {
			
			res[j] += res[j-1]
			
		}
		
	}
	return res
	
	//
	// res :=Triangles(n+1)
	// return res[n]
}
