package sort

// 基数排序

// 基数排序非常适合用于整数排序（尤其是非负整数）

// 执行流程 依次对个位数，十位数，百位数，千位数，万位数，进行排序（从低到高）
// 注意：不可以先从高位开始排序

func RadixSort(data []int) []int{
	
	var (
		counts = new([10]int) // 对每一位的0~9的数进行排序
		cle    = len(counts)
		le     = len(data)
		max    = maxVal(data)// 最大值 排序的次数
		tmp    = make([]int,le,le)// 每一次排好序的位置
		k int
	)
	
 	// 获取一共有多少个位数要排序 个,十,百...
	for c := 1; c <= max; c *= 10 {
		
		// 统计每个位数出现的次数
		for i := 0; i < le; i++ {
			counts[ data[i]/c%10 ]++
		}
		
		// 累加次数
		for j := 1; j < cle; j++ {
			counts[j] += counts[j-1]
		}
		
		// 从后往前便利元素，将它放到有序数组中的合适位置
		for n := le - 1; n >= 0; n-- {
			
			k = data[n] / c % 10
			tmp[counts[k]-1] = data[n]
			counts[k]-- // 相同数量位置-1
		}
		
		// 最终排序放回原来的数组，记录当前位数已有的序
		for i, v := range tmp {
			data[i] = v
		}
	}
	
	return data
}


// 获取数组的最大值
func maxVal(data []int) (max int) {
	
	max = data[0]
	
	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
		
	}
	
	return max
}