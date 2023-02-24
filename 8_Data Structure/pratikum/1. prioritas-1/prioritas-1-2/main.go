package main

import "fmt"

func Mapping(slice []string) map[string]int {
	res := map[string]int{}
	for _, v := range slice {
		res[v]++
	}
	return res
}

func main() {
	arr := []string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}
	fmt.Println(Mapping(arr))
}