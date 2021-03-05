package structure

import "fmt"

// heap
// 一种树状数据结构，一般用数据实现

// 数据顺序，从左到右，从上到下

// 索引i的规律
// 如果i = 0 它是根节点
// 如果i > 0 它的父节点编号等于 floor((i-1)/2)

//  n代表元素个数
// 如果 2i + 1 <= n-1 它的做子节点标号为2i+1
// 如果 2i + 1 > n-1，它无左子节点

// 如果 2i + 2 <= n-1  它的右子节点编号为2i+2
// 如果 2i + 2 > n - 1 它没有右子节点

// 堆结构
type MinHeap struct {
	Data []int
	Size int
}

// 清空堆
func (h *MinHeap) Clear() {
	h.Data = []int{}
	h.Size = 0
}

// 获取堆顶
func (h *MinHeap) Top() int {
	
	if h.Size > 0 {
		return h.Data[0]
	}
	
	panic("heap struct is empty")
}

// 是否是空堆
func (h *MinHeap) isEmpty() bool {
	return h.Size == 0
}

// Add(6)
// Add(16)
// Add(23)
// Add(18)
// Add(44)
// [6 16 23 18 44 ]

// 添加元素
func (h *MinHeap) Add(e int) {
	
	h.Data[h.Size] = e
	h.Size++
	h.SiftUp(h.Size - 1)
}

// 自上而下的上滤 让index位置元素上滤
// o(logn)
// 如果添加节点小于等于父节点直接退出
func (h *MinHeap) SiftUp(i int) {
	
	// @todo 方式1
	// for i > 0{
	//
	// 	// 当前指定节点
	// 	n := h.Data[i]
	//
	// 	// 父节点索引
	// 	pi:= (i - 1) >> 2
	// 	// 父节点值
	// 	pn:=h.Data[pi]
	//
	// 	// 添加节点小于父节点
	// 	if n <= pn {
	//
	// 		break
	// 	}
	//
	// 	// @todo 父节点与当前节点交换，每次都需要交换浪费资源
	// 	h.Data[pi], h.Data[i] = n, pn
	// 	// 重置当前索引到交换的父节点索引
	// 	i = pi
	// }
	//
	// 方式2 减少交换次数
	
	// 当前指定节点
	n := h.Data[i]
	
	for i > 0 {
		
		// 父节点索引
		pi := (i - 1) >> 2
		// 父节点值
		pn := h.Data[pi]
		
		// 添加节点小于父节点
		if n >= pn {
			break
		}
		
		// @todo 如果没找到添加节点的位置，证明添加节点比父节点大
		// 直接把父节点值放到当前节点上继续循环找父节点
		h.Data[i] = pn
		// 重置当前索引到交换的父节点索引
		i = pi
	}
	// 找到最后的位置
	h.Data[i] = n
}

// 删除 就是删除堆顶元素

// 1.删除元素时先拿最后一个元素覆盖他的位置
// 2.循环比较，做交换
// 如果node 小于子节点，与最大的子节点位置交换，或者没有字节点了 退出循环
func (h *MinHeap) Remove() int {
	
	if h.isEmpty() {
		panic("heap is empty")
	}
	
	// 减少元素
	h.Size--
	
	t := h.Data[0]
	// 覆盖元素
	h.Data[0] = h.Data[h.Size]
	// 删除尾元素
	h.Data = h.Data[:h.Size]
	
	if h.Size > 0{
		// 下滤
		h.SiftDown(0)
	}
	
	return t
}

// 自下而上的下滤
// 1.没有子节点退出循环
// 完全二叉树非叶子节点个数 n1+n2 = floor(n/2) 向下取整
// 先让本身节点和左右子节点建堆，然后继续下滤建堆

func (h *MinHeap) SiftDown(i int) {
	
	var (
		ci int // 用来比较的子节点索引
		cn int // 用来比较的子节点值
	)
	
	// 第一个个叶子节点的索引就是，非叶子节点的数量
	half := h.Size >> 1
	
	// 当前指定节点
	n := h.Data[i]
	
	// index位置保证是非叶子节点的位置才循环
	// 循环条件小于第一个叶子节点的索引
	for i < half {
		
		// 完全二叉树，非叶子节点有两种情况，1.只有左节点，2.左右都有
		
		// 默认用左子节点索引进行比较操作
		// 公式等于：2i+1
		li := (i << 1) + 1
		// 左子节点的值
		
		ln := h.Data[li]
		
		// 右子节点 index
		ri := li + 1
		
		// 判断是否存在右子节点,并且比较左右节点的值，谁比较大用谁
		if ri < h.Size && h.Data[ri] < ln {
			ci = ri
			cn = h.Data[ri]
		} else {
			ci = li
			cn = ln
		}
		
		// 如果添加节点大于子节点 直接退出
		if n <= cn {
			break
		}
		
		// 如果小于子节点那么用子节点替换当前节点
		h.Data[i] = cn
		
		// 重新赋值index位置用来下一次找到子节点子节点继续比较
		i = ci
	}
	
	h.Data[i] = n
	
}

// 删除堆顶元素的时候插入已个新元素
func (h *MinHeap) Replace(i int) {
	
	// @todo 方法1
	// 2*O(logn)
	// h.Remove()
	// h.Add(i)
	
	// @todo 优化版
	// 一次o(logn)
	if h.Size == 0 {
		h.Data[0] = i
		h.Size++
	} else {
		h.Data[0] = i
		h.SiftDown(0)
	}
	
}

// 建堆
// 时间复杂度 O(nlogn)
// 1.自上而下的上滤(堆顶不需要上滤)，从索引1开始上滤
//  >> 相当于每个节点都做上滤操作
//  for(i:=1;i<size;i++){
//      SiftUp(i)
//  }
// 时间复杂度 O(n)
// 2.自下而上的下滤,排除叶子节点 (size >>1 )-1 最后一个非叶子节点索引
//  >> 效率比较高，因为下滤操作节点比较少，且只有在顶层的元素才需要这个操作，如果不在顶层下滤操作会很少
//  for(i:=(size >>1 )-1;i>=0;i--){
//       SiftDown(i)
//  }
func (h *MinHeap) HeapIfy(Data []int) {
	
	// @todo 这种方法是浅copy 会导致外面修改了内部的值
	// h.Data =Data
	// h.Size =len(Data)
	
	l := len(Data)
	
	if h.Data == nil || h.Size == 0 {
		h.Data = make([]int, l)
	}
	
	// 添加到数组里面
	for i, v := range Data {
		h.Data[i] = v
	}
	
	h.Size = l
	
	fmt.Println(h.Data)
	// 堆化操作
	for i := (h.Size >> 1) - 1; i >= 0; i-- {
		h.SiftDown(i)
	}
}

// 疑惑:
// 以下方式是否可以建堆
// 1.自上而下的下滤
// 2.自下而上的上滤

// 不可以
