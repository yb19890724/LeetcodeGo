package sort

// 希尔排序 别名 递减增量排序
// 希尔排序底层一般使用插入排序每一列进行排序，(很多资料认为希尔排序是插入排序改良版)

// 把序列看成是一个"矩阵"，分成m列，逐列进行排序
// m从某个整数逐渐减为1
// 当m为1，整个序列将完全有序

// 矩阵的列数取决于步长序列 (step sequence)
// 例： 如果步长序列为{1,5,19,41,109,...} 就代表依次分成109列，41列，19列，5列，1列
// 不同的步长序列，执行效率也不同

// @todo 步长序列理解
// n/2k 比如n为16(数据规模) k次方，
// 2k(1) = 16/2  == 8
// 2k(2) = 16/4  == 4
// 2k(3) = 16/8  == 2
// 2k(4) = 16/16 == 1
// 步长序列是{1,2,4,8}

// @todo 过程分析
// 数据规模  步长{5,2,1}
// 11,10,9,8,7,6,5,4,3,2,1

// 排序列              排序后
// <11>,10,9,8,7     <1>,10,9,8,7
// <6>,5,4,3,2       <6>,5,4,3,2
// <1>               <11>,5,4,3,2

// 假设元素在第col列，第row行，步长(总列数) 是step
// 那么这个元素在数组中的索引是 col+row * step
// 例子

// 9 排序前是第2列，第0行，排序索引 2+0*5=2
// 1 排序前是第2列，第1行，排序索引 2+1*5=7

func ShellSort(data []int) []int {
	
	le := len(data)
	// 构建步长序列
	steps := stepSequence(data)
	
	for _,step := range steps {
		// 开始插入排序，每一轮的步长为 step
		
		// 每次循环对应步长增加
		for i := step; i < le; i += step {
			//  index 5 比较 index 0
			for j := i - step; j >= 0; j -= step {
				
				// 满足插入那么交换元素
				if data[j+step] < data[j] {
					data[j], data[j+step] = data[j+step], data[j]
					continue
				}
				break
			}
		}
	}
	return data
}

// 步长序列
func stepSequence(data[]int) []int {
	
	var (
		s []int
		n = len(data)
	)
	
	//for step := n / 2; step >= 1; step /= 2 {
	for step := n >> 1; step >= 1; step >>= 1 {
		s = append(s,step)
	}
	
	return s
}