package main

import "fmt"

// 默认棋局大小
// var chessMap [][]int

// 值节点
type ValNode struct {
	row int
	col int
	val int
}

// 打印棋局
func ShowChessMap(chessMap [11][11]int) {
	// 输出看原始数组
	
	for _, v := range chessMap {
		
		for _, v2 := range v {
			
			fmt.Printf("%d\t", v2)
		}
		
		fmt.Println()
	}
	
}

// 存储棋局
func SaveChessMap(chessMap [11][11]int) []ValNode {
	
	// 创建一个值节点
	var sparseArr []ValNode
	
	// 标准稀疏数组应该有，原始数据的规模（行和列，默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	
	// 存储稀疏数组值
	for i, v := range chessMap {
		
		for j, v2 := range v {
			
			if v2 != 0 {
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
		
	}
	return sparseArr
}

// 恢复棋盘
func LoadGame(sparseArr[]ValNode) [11][11]int{
	
	// 存盘功能可以用 写入文件方式
	// 回复用读取文件方式
	
	sparseArr = sparseArr[1:] // 去掉棋盘的大小
	
	var chessMap [11][11]int
	// 便利 每一行文件
	for i, valNode := range sparseArr {
		
		if i != 0 { // 这里判断是因为兼容 第一个存储的是规模
			
			chessMap[valNode.row][valNode.col] = valNode.val
		}
		
	}
	return chessMap
}

func main() {
	
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑棋子
	chessMap[2][3] = 2 // 白棋子
	
	// 查看棋局
	ShowChessMap(chessMap)
	
	// 存储棋局和棋盘
	sparseArr := SaveChessMap(chessMap)
	
	// 查看存储的棋局
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	
	// 读取进度
	LoadGame(sparseArr)
	
	
	// 读取棋盘进度
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
