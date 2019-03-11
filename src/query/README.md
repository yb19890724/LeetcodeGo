#### 二分查找

##### 普通版本




##### 二分查找第一匹配的元素

data[middle]跟要查找的value的大小关系有三种情况:大于、小于、等于。
对于data[middle]>value的情况，我们需要更新high= mid-1;
对于data[middle]<value的情况，我们需要 更新low=mid+1。
那当data[middle]=value的时候应该如何处理呢?
当data[middle]等于要查找的值时,data[middle]就是我们要找的元素。
但是，如果我们求解的是第一个值等于给定值的元素，当data[mid]等于要查找的值时，
我们就需要确认一下这个data[mid]是不是第一个值等于给定值的元素。
如果mid等于0，那这个元素已经是数组的第一个元素，那它肯定是我们要找的;
如果middle不等于0，但data[mid]的前一个元素data[mid-1]不等于value，那也说明data[mid]就是我们要找的第一个值等于给定值的元素。



