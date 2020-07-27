package sort

// 插入排序

// 插入排序类似扑克牌 排序

// 执行流程
// 1.在执行过程中，插入排序会将序列分为2部分
// 头部已经排好序的，尾部 是待排序的

// 2.从头开始扫描每一个元素
// 每当扫描到一个元素 ，就将它插入到头部适合的位置，使得头部数据依然保持有序

func InsertSort(data []int) []int {
	
	for i := range data {
		
		preIndex := i - 1
		current := data[i]
		
		for preIndex >= 0 && data[preIndex] > current {
			data[preIndex+1] = data[preIndex]
			preIndex -= 1
		}
		data[preIndex+1] = current
	}
	return data
	
	/*
	for i := 0; i < len(data); i++ {
		
		deal := data[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置
		
		// 如果第一次比较，比左边的已排好序的第一个数小，那么进入处理
		if deal < data[j] {
			// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
			for ; j >= 0 && deal < data[j]; j-- {
				data[j+1] = data[j] // 某数后移，给待排序留空位
			}
			data[j+1] = deal // 结束了，待排序的数插入空位
		}
		
	}*/
}