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
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}