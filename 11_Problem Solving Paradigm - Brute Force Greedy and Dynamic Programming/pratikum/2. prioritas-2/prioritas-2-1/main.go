package main

import "fmt"

func Frog(jumps []int) int {
	n := len(jumps) - 1

	if n <= 2 {
		return getAbs(jumps[0] - jumps[n])
	}
	costs := make([]int, n+1)
	costs[0] = 0
	costs[1] = getAbs(jumps[0] - jumps[1])

	for i := 2; i <= n; i++ {
		prevJump1 := costs[i-1] + getAbs(jumps[i]-jumps[i-1])
		prevJump2 := costs[i-2] + getAbs(jumps[i]-jumps[i-2])
		costs[i] = getMin(prevJump1, prevJump2)
	}

	return costs[n]
}


func getAbs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}