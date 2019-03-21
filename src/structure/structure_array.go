package structure

// 值节点
type ValNode struct {
	Row int // 记录行
	Col int // 记录列
	Val int // 记录值
}

// 创建记录数据范围值

func CreateArray(row int, col int) [][]int {
	var data = make([][]int, row, col)
	
	for i := 0; i < col; i++ {
		var rowArray []int = make([]int, row, col)
		data[i] = rowArray
	}
	return data
}

// 恢复原始数据
// @var sparseArr[0]  	记录行，记录记录列
// @var sparseArr[0:]   记录值

func RecoverSparseArr(sparseArr []ValNode) [][]int {
	
	size, sparseArr := sparseArr[0], sparseArr[0:]
	
	data := CreateArray(size.Row, size.Col)
	
	// 便利 每一行文件
	for i, valNode := range sparseArr {
		
		if i != 0 { // 这里判断是因为兼容 第一个存储的是规模
			
			data[valNode.Row][valNode.Col] = valNode.Val
		}
		
	}
	return data
}

// 存储记录到稀疏数组
func SaveSparseArr(data [][]int, sparseArr []ValNode) []ValNode {
	
	// 存储稀疏数组值
	for i, v := range data {
		
		for j, v2 := range v {
			
			if v2 != 0 {
				var valNode = ValNode{
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
		
	}
	return sparseArr
}
