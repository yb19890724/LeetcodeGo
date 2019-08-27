package sort

// 计数排序
/*
   @data	  []int{2,5,3,0,2,3,0,3}
   @for  	  计算数组中最大值
   @dataTemp  申请一个临时最大值加1长度的数组,用来统计数出现的次数,因为有0存在所以加1
   		 	  循环计算每一位出现的次数,如果出现了就+1,通过遍历计算临时存储数组值放入到最终数据
			  []int{2,0,2,3,0,1} 例对应：0有2个,1有0个..
   @dataTemp  把数组值从左往右以此相加  []int{2, 2 ,4 ,7 ,7 ,8}
			  注：2代表小于等于0的数据，4小于等于2的数
   @result	  []int{2,5,3,0,2,3,0,3}  []int{2, 2 ,4 ,7 ,7 ,8}
			  dataTemp[data[i]] - 1 通过原始数据值，找到桶中对应值得位置,-1是因为数组从0开始
			  result[index] = data[i] 根据位置存放取出的值
			  dataTemp[data[i]]--	桶中被取走了一个值，所以位置减去一个
				例：存在相同值 3和3的排序
*/

func CountSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	var max int = data[0]

	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	// @todo 分配桶
	var dataTemp []int = make([]int, max+1)

	for i := 0; i < len(data); i++ {
		dataTemp[data[i]]++
	}

	for i := 1; i <= max; i++ {
		dataTemp[i] = dataTemp[i-1] + dataTemp[i]
	}

	var result []int = make([]int, len(data))

	for i := len(data) - 1; i >= 0; i-- {

		index := dataTemp[data[i]] - 1

		result[index] = data[i]

		dataTemp[data[i]]--

	}

	return result
}
