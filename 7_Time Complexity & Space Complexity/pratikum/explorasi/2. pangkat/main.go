package main

import "fmt"

func pow(x, n int) int {
	
	res := 1
	for i:= n; i > 0; i /=2 {
		if i % 2 != 0 {
			res *= x
		}
		x *=x
	}
	
	return res

}

func main() {
	fmt.Println(pow(2,8))

}