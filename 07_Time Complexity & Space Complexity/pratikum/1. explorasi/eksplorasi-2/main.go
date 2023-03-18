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
	fmt.Println(pow(2,3))
	fmt.Println(pow(5,3))
	fmt.Println(pow(10,2))
	fmt.Println(pow(2,5))
	fmt.Println(pow(7,3))

}