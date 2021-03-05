package sort


// 计数排序

// 统计每个整数在序列中出现的次数，进而推导出每个整数在有序序列中的索引

// 推导分析:
// 数据源：[]int{7,3,5,8,6,7,4,5}

// 存放index 0,1,2,3,4,5,6,7,8
// 存放次数   0,0,0,1,1,2,1,2,1

// @todo 简单缺点
// 1.非常浪费空间
// 2.无法对负数排序
// 3.不稳定排序

// 时间复杂度 3*O(n) = O(n)

func CountSort(data[]int) []int {
	
	le := len(data)
	
	if le <= 1 {
		return data
	}
	
	var (
		max       = data[0] // 最大值
		sortIndex = 0       // 排序开始
	)
	
	for i := 1; i < le; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	
	// 记录整数出现的次数
	buckets := make([]int,max+1)

	for _,v:=range data {
		buckets[v]+=1
	}
	
	// 根据帧数出现次数，对整数进行排序
	
	ble := len(buckets)
	for j := 0; j < ble; j++ {
		for buckets[j] > 0 {
			data[sortIndex] = j
			sortIndex += 1
			buckets[j] -= 1
		}
	}
	
	return data
}

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

func CountSort1(data []int) []int {
	
	if len(data) <= 1 {
		return data
	}
	
	var max = data[0]
	
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	
	// @todo 分配桶
	var dataTemp = make([]int, max+1)
	
	for i := 0; i < len(data); i++ {
		dataTemp[data[i]]++
	}
	
	// 每次次数加上前面的所有次数,这样可以计算前面有多少个数
	for i := 1; i <= max; i++ {
		dataTemp[i] = dataTemp[i-1] + dataTemp[i]
	}
	
	var result = make([]int, len(data))
	
	for i := len(data) - 1; i >= 0; i-- {
		
		index := dataTemp[data[i]] - 1
		
		result[index] = data[i]
		
		dataTemp[data[i]]--
		
	}
	
	return result
}


// min data里面的最小值
// index = k - min
// array 元素k 在有序序列中的元素索引
// counts [k-min] - p
// p代表着是倒数第几个k

// min = 3
// 元素       3      4        5       6      7       8
// 索引   0=(3-3) 1=(4-3) 2=(5-3) 3=(6-3) 4=(7-3) 5=(8-3)

// 优化空间

func CountSort2(data []int) []int {
	
	if len(data) <= 1 {
		return data
	}
	
	var max = data[0]
	var min = data[0]
	
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		
		if data[i] > min {
			min = data[i]
		}
	}
	
	// @todo 分配桶
	var dataTemp = make([]int, max+1)
	
	for i := 0; i < len(data); i++ {
		dataTemp[data[i]]++
	}
	
	// 每次次数加上前面的所有次数,这样可以计算前面有多少个数
	for i := 1; i <= max; i++ {
		dataTemp[i] = dataTemp[i-1] + dataTemp[i]
	}
	
	var result = make([]int, len(data))
	
	for i := len(data) - 1; i >= 0; i-- {
		
		index := dataTemp[data[i]] - 1
		
		result[index] = data[i]
		
		dataTemp[data[i]]--
		
	}
	
	return result
}