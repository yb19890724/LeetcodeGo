package sort

// 桶排序

// 1.创建一定数量的桶（比如数组或者链表）
// 2.按照一定的规则(不同类型的数据，规则不同)，将序列中的元素均匀分配大盘对应的桶
// 3.分别对每个桶进行单独排序
// 4.将所有非空桶的元素合并成有序序列

// 空间复杂度: O(n+m) m是桶的数量(m是有数据而定的)

func sortInBucket(bucket []int) { // 此处实现插入排序方式，其实可以用任意其他排序方式
	length := len(bucket)
	if length == 1 {
		return
	}
	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1;
		// 将选出的被排数比较后插入左边有序区
		for j >= 0 && backup < bucket[j] { // 注意j >= 0必须在前边，否则会数组越界
			bucket[j+1] = bucket[j] // 移动有序数组
			j--                     // 反向移动下标
		}
		bucket[j+1] = backup // 插队插入移动后的空位
	}
}


// 桶排序
func BucketSort(data []int) []int {
	// 桶数
	num := len(data)
	// k（数组最大值）
	max := maxVal(data)
	// 二维切片
	buckets := make([][]int, num)
	// 分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = data[i] * (num - 1) / max // 分配桶index = value * (n-1) /k
		buckets[index] = append(buckets[index], data[i])
	}
	
	// 桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])
			copy(data[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return data
}
