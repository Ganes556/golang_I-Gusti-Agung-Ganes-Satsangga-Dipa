package main

import "fmt"

func selisih(matriks [][]int) int {
	diagonalKiri := 0
	diagonalKanan := 0
	for i, v := range matriks {
		diagonalKiri += v[i]
		diagonalKanan += v[len(v)-i-1]
	}
	selisih := diagonalKanan - diagonalKiri
	if selisih < 0 {
		selisih *= -1
	}
	return selisih
}

func main() {
	fmt.Println(selisih([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}