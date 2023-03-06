package main

import "fmt"

func fibo(n int, memo []int) int {
	if n < 2 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibo(n-1, memo) + fibo(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println(fibo(0,make([]int, 1)))
	fmt.Println(fibo(1,make([]int, 2)))
	fmt.Println(fibo(2,make([]int, 3)))
	fmt.Println(fibo(3,make([]int, 4)))
	fmt.Println(fibo(4,make([]int, 5)))
	fmt.Println(fibo(5,make([]int, 6)))
	fmt.Println(fibo(6,make([]int, 7)))
	fmt.Println(fibo(7,make([]int, 8)))
	fmt.Println(fibo(8,make([]int, 9)))
	fmt.Println(fibo(9,make([]int, 10)))
	fmt.Println(fibo(10,make([]int, 11)))
}