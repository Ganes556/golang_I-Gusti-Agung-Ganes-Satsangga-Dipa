package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	mergeDict := map[string]int{}
	mergeSlice := []string{}
	for _,v := range append(arrayA,arrayB...) {
		mergeDict[v]++
		if val, ok := mergeDict[v]; ok == true && val == 1{
			mergeSlice = append(mergeSlice, v)
		}
	}
	return mergeSlice

}
func main() {
	fmt.Println(ArrayMerge([]string{"king","devil jin","akuma"},[]string{"eddie","steve","geese"}))

	fmt.Println(ArrayMerge([]string{"sergei","jin"},[]string{"jin","steve","bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa","yoshimitsu"},[]string{"devil jin","yoshimitsu","alisa","law"}))

	fmt.Println(ArrayMerge([]string{},[]string{"devil jin","sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"},[]string{}))

	fmt.Println(ArrayMerge([]string{},[]string{}))
	
}