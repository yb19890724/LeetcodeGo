package sort

// 计数排序
/*
   计算数组中最大值
   申请一个临时最大值加1长度的数组,用来统计数出现的次数,因为有0存在所以加1
   循环计算每一位出现的次数,如果出现了就+1
   通过遍历计算临时存储数组值放入到最终数据
*/

func CountSort(data []int) []int {
	
	if nil == data || len(data) <= 0 {
		return data
	}
	
	var max int = data[0]
	
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	
	var dataTemp []int = make([]int, len(data)+1)
	
	for i := 0; i < len(data); i++ {
		dataTemp[data[i]]++
	}
	
	var r []int =make([]int,max+1)
	for i := len(data)-1; i >=0; i-- {
		index :=dataTemp[data[i]]-1
		r[index] = data[i]
		dataTemp[data[i]]--
	}
	
	for i := 0; i < len(data); i++ {
		data[i] = r[i]
	}
	
	return data
}
