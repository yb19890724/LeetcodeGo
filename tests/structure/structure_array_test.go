package structure

import (
	"fmt"
	"github.com/yb19890724/leetcode-go/src/structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 打印棋局
func ShowChessMap(data [][]int) {
	// 输出看原始数组
	
	for _, v := range data {
		
		for _, v2 := range v {
			
			fmt.Printf("%d\t", v2)
		}
		
		fmt.Println()
	}
	
}

// @todo: 创建记录行测试数据

func createTestData(row int, col int) ([]int, []int, []int) {
	
	var defaultRow = make([]int, row, col)
	
	var twoRow, threeRow = make([]int, row, col), make([]int, row, col)
	
	twoRow[2], threeRow[3] = 1, 2
	
	return defaultRow, twoRow, threeRow
}

// @todo 创建棋局数据，模拟棋局数据

func createChessMapData(row int, col int) [][]int {
	
	chessMap := createArray(row, col)
	
	chessMap[1][2], chessMap[2][3] = 1, 2
	
	return chessMap
}

// @todo 创建数组
func createArray(row int, col int) [][]int {
	
	chessMap := structure.CreateArray(row, col)
	
	return chessMap
}

// @todo 创建稀疏数组记录
func createSparseArray(row int, col int) []structure.ValNode {
	
	sparseArray := []structure.ValNode{
		{
			Row: row,
			Col: col,
			Val: 0,
		},
		{
			Row: 1,
			Col: 2,
			Val: 1,
		},
		{
			Row: 2,
			Col: 3,
			Val: 2,
		},
	}
	return sparseArray
	
}

// @test create sparse array
func TestCreateSparseArray(t *testing.T) {
	
	row, col := 11, 11 // 定义棋局范围
	
	chessMap := createChessMapData(row, col)
	
	defaultRow, twoRow, threeRow := createTestData(row, col)
	
	for index, v := range chessMap {
		
		if 1 == index {
			
			assert.Equal(t, v, twoRow, " default row comparative failure ")
			continue
		}
		
		if 2 == index {
			
			assert.Equal(t, v, threeRow, " two row comparative failure")
			continue
		}
		
		assert.Equal(t, v, defaultRow, " two row comparative failure")
	}
	
}

// @test save sparse array
func TestSaveSparseArray(t *testing.T) {
	
	row, col := 11, 11 // 定义棋局范围
	
	chessMap := createChessMapData(row, col)
	
	var sparseArr []structure.ValNode
	
	valNode := structure.ValNode{
		Row: row,
		Col: col,
		Val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	
	saveChessMap := structure.SaveSparseArray(chessMap, sparseArr)
	
	actualChessMap := createSparseArray(row, col)
	
	assert.Equal(t, actualChessMap, saveChessMap, " save sparse array error ")
	
}

// @test  recover sparse array
func TestRecoverSparseArray(t *testing.T) {
	
	row, col := 11, 11 // 定义棋局范围
	
	recodeChessMap := createSparseArray(row, col)
	
	actualChessMap := structure.RecoverSparseArray(recodeChessMap)
	
	defaultRow, twoRow, threeRow := createTestData(row, col)
	
	for index, v := range actualChessMap {
		
		if 1 == index {
			
			assert.Equal(t, v, twoRow, " default row comparative failure ")
			continue
		}
		
		if 2 == index {
			
			assert.Equal(t, v, threeRow, " two row comparative failure")
			continue
		}
		
		assert.Equal(t, v, defaultRow, " two row comparative failure")
		
	}
}
