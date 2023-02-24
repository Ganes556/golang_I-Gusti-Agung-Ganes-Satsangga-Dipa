package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	return append(arrayA,arrayB...)
}
func main() {
	arr1 := []string{"broh","broh2"}
	// arr2 := []string{"broh3","broh4"}

	fmt.Println(arr1)
}