package structure


// 杨辉三角

func Triangles(n int) [][]int {
	
	res := CreateArray(n,n)
	
	for i := 0; i < n; i++ {
		
		for j := 0; j <= i; j++ {
			
			if j == 0 || j == i {
				res [i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
			//fmt.Print(res[i][j])
			
		}
		
		//fmt.Println("")
	}
	return  res
}
