package main

import (
	`fmt`
)

// 默认棋局大小
// var chessMap [][]int

// 值节点
type ValNode struct {
	row int
	col int
	val int
}

// 打印棋局
func ShowChessMap(chessMap [][]int) {
	// 输出看原始数组
	
	for _, v := range chessMap {
		
		for _, v2 := range v {
			
			fmt.Printf("%d\t", v2)
		}
		
		fmt.Println()
	}
	
}

// 恢复棋盘
func LoadGame(sparseArr []ValNode) [][]int {
	
	chessboardLayout := sparseArr[0]  // 获取棋盘的范围
	sparseArr = sparseArr[0:] // 去掉棋盘的大小
	
	chessMap := CreateChessMap(chessboardLayout.row, chessboardLayout.col)
	
	// 便利 每一行文件
	for i, valNode := range sparseArr {
		
		if i != 0 { // 这里判断是因为兼容 第一个存储的是规模
			
			chessMap[valNode.row][valNode.col] = valNode.val
		}
		
	}
	return chessMap
}

// 创建空棋局
func CreateChessMap(row int, col int) [][]int {
	var chessMap = make([][]int, row, col)
	
	for i := 0; i < col; i++ {
		var rowArray []int = make([]int, row, col)
		chessMap[i] = rowArray
	}
	return chessMap
}

// 存储棋局
func SaveChessMap(chessMap [][]int, sparseArr []ValNode) []ValNode {
	
	// 存储稀疏数组值
	for i, v := range chessMap {
		
		for j, v2 := range v {
			
			if v2 != 0 {
				var valNode = ValNode{
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

func main() {
	
	row, col := 11, 11 // 定义棋局范围
	chessMap := CreateChessMap(row, col)
	
	chessMap[1][2] = 1 // 黑棋子
	chessMap[2][3] = 2 // 白棋子
	
	ShowChessMap(chessMap)
	
	// 创建一个值节点
	var sparseArr []ValNode
	
	// 标准稀疏数组应该有，原始数据的规模（行和列，默认值）
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	
	// 存储棋局和棋盘
	saveChessMap := SaveChessMap(chessMap, sparseArr)
	
	
	// 查看存储的棋局
	for i, valNode := range saveChessMap {
	 	fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	
	// 读取进度
	loadGame := LoadGame(saveChessMap)
	
	// 读取棋盘进度
	for _, v := range loadGame {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
