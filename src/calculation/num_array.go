package calculation

// 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
//
// 示例：
//
// 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
//
// sumRange(0, 2) -> 1
// sumRange(2, 5) -> -1
// sumRange(0, 5) -> -3
//
// 链接：https://leetcode-cn.com/problems/range-sum-query-immutable

type NumArray struct {
	data []int
}


func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

// @todo 暴力算法
// 暴力法的最大缺点就是耗时，像这里的 sumRange 方法中有一层for循环，时间复杂度为 O(n)，
// 多次调用 sumRange 方法，所以调用 k 次的话，时间复杂度则为 O(kn) 。
func (n *NumArray) SumRange(i int, j int) int {
	sum := 0
	for k:= i; k <=j; k++ {
		sum += n.data[k]
	}
	return sum
}



// @todo 优化 缓存方式
// 我们可以提前计算出前 i 个数的和，即 a[i] 等于区间 [1,i] 的和，消耗时间为 O(n)
//
// 区间 [i,j] 的和求公式为：SumRange(i,j) = sum[j] - sum[i-1]，每次查询消耗时间为 O(1)，k 次查询时间复杂度为 O(k)
//
// 总的来说，时间复杂度为 O(n+k)

func Constructor2(nums []int) NumArray {
	a := NumArray{}
	a.data = append(a.data, 0) // 初始化 data 切片，data[0]=0
	// 遍历求sum[i]
	for i := 1; i <= len(nums); i++ {
		// 当前节点存储的是 前一个节点加上
		a.data = append(a.data, a.data[i-1]+nums[i-1])
	}
	return a
}

// 直接找到 start 和 end 的和 end - start = 范围和
func (n *NumArray) SumRange2(i int, j int) int {
	return n.data[j+1] - n.data[i]
}











