package query

var (
	low    int // 开始区间
	high   int // 结束区间
	middle int // 待查找区间的中间元素下标
)

// 二分查找
// 在有序数组中查找int数值
// ts: O(log*n)

/*
  @todo:理解
  实际上，mid=(low+high)/2这种写法是有问题的。
  因为如果low和high比较大的话，两者之和就有可能会溢出。
  改进的方法是将mid的计算方式写成low+(high-low)/2。
  更进一步，如果要将性能优化到极致的话，我们可以将这里的除以2操作转化成位运算low+((high-low)>>1)
*/
func BinarySort(data []int, search int) bool {
	
	low = 0
	high = len(data) - 1
	
	for low <= high {
		middle = low + (high-low)/2
		
		if data[middle] > search {
			high = middle - 1
		} else if data[middle] < search {
			low = middle + 1
		} else {
			return true
		}
	}
	return false
}

// 查找第一匹配的元素

/*
  @todo:理解
  data[middle]跟要查找的value的大小关系有三种情况:大于、小于、等于。
  对于data[middle]>value的情况，我们需要更新high=middle-1;
  对于data[middle]<value的情况，我们需要 更新low=middle+1。
  那当data[middle]=value的时候应该如何处理呢?
  当data[middle]等于要查找的值时,data[middle]就是我们要找的元素。
  但是，如果我们求解的是第一个值等于给定值的元素，当data[middle]等于要查找的值时，
  我们就需要确认一下这个data[middle]是不是第一个值等于给定值的元素。
  如果mid等于0，那这个元素已经是数组的第一个元素，那它肯定是我们要找的;
  如果middle不等于0，但data[middle]的前一个元素data[middle-1]不等于value，那也说明data[middle]就是我们要找的第一个值等于给定值的元素。
  如果经过检查之后发现data[middle]前面的一个元素data[middle-1]也等于value，那说明此时的data[middle]肯定不是,我们要查找的第一个值等于给定值的元素。
  那我们就更新high=middle-1，因为要找的元素肯定出现 在[low, middle-1]之间。
*/

func FirstBinarySearch(data []int, search int) bool {
	
	low = 0
	high = len(data) - 1
	
	for low <= high {
		middle = low + (high-low)/2
		
		if data[middle] > search {
			high = middle - 1
		} else if data[middle] < search {
			low = middle + 1
		} else {
			if (middle == 0) || (data[middle-1] != search) {
				return true
			}
			high = middle - 1
		}
	}
	return false
}

// 查找最后一个匹配的元素
/**
	如果data[middle]这个元素已经是数组中的最后一个元素了，那它肯定是我们要找的；
	如果data[middle]的后一个元素data[middle+1]不等于value，那也说明data[middle]就是我们要找的最后一个值等于给定值的元素。
 	发现data[middle]后面的一个元素data[middle+1]也等于value，那说明当前的这个data[middle]并不是最后一个值等于给定值的元素。
	我们就更新low=middle+1，因为要找的元素肯定出现在[middle+1, high]之间。
*/

func LastBinarySearch(data []int, search int) bool {
	low = 0
	high = len(data) - 1
	
	for low <= high {
		middle = low + (high-low)/2
		
		if data[middle] > search {
			high = middle - 1
		} else if data[middle] < search {
			low = middle + 1
		} else {
			if (middle == len(data)-1) || (data[middle+1] != search) {
				return true
			}
			low = middle + 1
		}
	}
	return false
}
