package main

import "fmt"

func mergeInt(integers ...int) (integer int) {
	length := len(integers)
	integer = integers[length-1]
	for i:= len(integers)-2; i >=0 ; i--{
		integer = integer*10 + integers[i]
	}
	return integer
}


func binerX(n int) {
	binTable := make([]int, 0, n+1)
	binTable = append(binTable, 0)
	for i := 1; i <= n; i++ {
		bin := []int{}
		for j:= i; j > 0; j/=2{
			bin = append(bin, j%2)
		}
		if len(bin) > 0 {
			binTable = append(binTable, mergeInt(bin...))
		}
	}
	fmt.Println(binTable)
}

func main() {
	binerX(2)
	binerX(3)
}