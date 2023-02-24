package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	res := []int{}
	for i, v := range angka {
		r := v
		for j, v2 := range angka {
			if i == j {
				continue
			}
			if v == v2 {
				r = 0
				break
			}
					
		}
		if k, err := strconv.Atoi(string(r)); err == nil && r != 0{
			res = append(res, k)
		}
	}
	return res
}

func main() {
	fmt.Println(munculSekali("1234566789"))
}