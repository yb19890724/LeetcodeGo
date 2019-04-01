package interviews

import (
	"fmt"
	"strings"
)

//@test multiplication table
func MultiplicationTable() []string {

	var result []string

	for i := 1; i <= 9; i++ {

		var formula string

		for j := 1; j <= i; j++ {

			item := fmt.Sprintf("%d x %d = %d  ", i, j, i*j)

			formula = strings.Join([]string{formula, item}, "")

		}

		result = append(result, formula)
	}
	return result
}
