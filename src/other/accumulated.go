package other

// 有一组不同高度的台阶，有一个整数数组表示，数组中每个数是台阶的高度，当开始下雨了（雨水足够多）台阶之间的水坑会积水多少呢？
// 如下图，可以表示为数组[0,1,0,2,1,0,1,3,2,1,2,1] 返回


//https://blog.csdn.net/Thomas0713/article/details/83051990
// 思路：
// 1.找到最大值的索引位
// 2.定义两个判方向，
//      1）从最大值向右，当前点位右面的值大于当前点位则可以积水
//      2）从最大值向左, 当前点位左面的值大于当前点位则可以积水
func WaterVolume(accurmulated []int) int{
	
	arrMax := 0                 // 定义数组的最大值
	arrMaxPos := 0              // 定义数组最大值的下标
	
	for index, v := range accurmulated {
		
		if arrMax < v {
			arrMax = v
			arrMaxPos = index
		}
		
	}
	
	arrMaxLeft :=0 // 从左开始遍历 极大值
	arrMAxRight:=0 // 定义从最右边遍历极大值
	volumes :=0 // 定义水容量
	
	for i:=0;i<arrMaxPos;i++  {
		
		if arrMaxLeft < accurmulated[i]{
			arrMaxLeft= accurmulated[i]
		}else{
			volumes+=(arrMaxLeft- accurmulated[i])
		}
	}

	
	for i:=len(accurmulated)-1;i>arrMaxPos;i--  {
		if arrMAxRight < accurmulated[i]{
			arrMAxRight= accurmulated[i]
		}else{
			volumes+=(arrMAxRight- accurmulated[i])
		}
	}
	return volumes
}
