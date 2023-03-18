package main

import "fmt"

// Dynamic Programming, Method Bottom up with tabulation
func rowSegitigaPascal(row int) [][]int {
	if row <= 0 {
		return [][]int{{}}
	}
	if row == 1 {
		return [][]int{{1}}
	}
	
	pascalRows := make([][]int, row)
	pascalRows[0] = []int{1}
	pascalRows[1] = []int{1,1}

	for i:= 2; i < row; i++ {
		row := make([]int,i+1)
		row[0] = 1
		row[i] = 1
		for j:= 1; j < i; j++ {
			row[j] = pascalRows[i-1][j-1] + pascalRows[i-1][j]
		}
		pascalRows[i] = row
	}

	return pascalRows
}

func main() {
	fmt.Println(rowSegitigaPascal(5))
}