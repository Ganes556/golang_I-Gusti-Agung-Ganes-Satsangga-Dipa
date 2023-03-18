package main

import "fmt"

func PairSum(arr []int, target int) []int {
	n := len(arr) - 1
	
	res := []int{}
	loop:
	for n >= 1 {
		// check mid data
		if arr[(n/2)+1] >= target {
			n /= 2
			continue
		}
		// check last data
		if arr[n] == target {
			n -= 1
			continue
		}
		// check rest data from the last until the first
		for i := n - 1; i >= 0; i-- {
			if arr[i]+arr[n] == target {
				res = append(res, []int{i, n}...)
				break loop
			}
		}
		n -= 1
	}
	return res
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 10}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
	fmt.Println(PairSum([]int{1, 4, 9, 10, 12, 90, 100, 101,102}, 100))
}