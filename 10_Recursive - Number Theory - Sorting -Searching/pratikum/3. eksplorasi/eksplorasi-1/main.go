package main

import (
	"fmt"
)

func getMax(sum, crrNum int) int{
	if sum > crrNum {
		return sum
	}
	return crrNum
}

func MaxSequence(arr []int) int{
	// your code here

	// menggunakan algoritma Kadane's
	sum := arr[0]
	currentMax := arr[0]
	for i := 1; i < len(arr); i++{
		sum = getMax(sum + arr[i], arr[i])
		currentMax = getMax(sum, currentMax)
	}
	
	return currentMax

}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))  // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))    // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))    // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))    // 8
  fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))     // 12
}