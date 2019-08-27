package interviews

import "fmt"

// 打印金字塔
// @param level 层级
func Pyramid(level int) {

	for i := 1; i <= level; i++ {

		for j := 1; j <= level-i; j++ {
			fmt.Print(" ")
		}
		//k表示每层打印多少*
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

}
