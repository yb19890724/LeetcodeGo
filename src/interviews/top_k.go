package interviews

import "github.com/yb19890724/leetcode-go/src/structure"

// 从n个整数，找出最大的前k个数(k 远远小于n)

// 解法： k=10举例

// 100w数排序 n和k比较  10远远小于100w
// 1. 快速排序算法，需要O(nlogn)的时间复杂度
// 2. 二叉堆，可以使用O(nlogk)的时间复杂度解决


// 堆解决：小顶堆
// 1.扫描n个整数 100w数全部需要扫描
// 2.先遍历到前k个数放入堆中
// 3.从k+1个数开始，如果大于堆顶元素，就使用replace操作(删除堆顶元素,将第k+1个数添加到堆中)

// 最大前k个数
var (
	k    = 5
	data = []int{1, 10, 230, 4, 99, 21, 334, 20, 60, 50}
	l    = len(data)
)


func TopK() {
	var h structure.MaxHeap
	
	for i := 0; i < l; i++ {
		
		// 小于前k个数
		if h.Size < k {
			
			h.Add(data[i]) // O(logk)
			continue
		}
		
		if data[i] > h.Data[0] {// 如果是k+1 替换操作
			// 如果大于小顶堆的堆顶 就替换
			h.Replace(data[i])// O(logk)
		}
	}
}

