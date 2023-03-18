package main

import "fmt"

func MergeSortDesc(cards [][]int) [][]int{
	if len(cards) <= 1 {
		return cards
	}
	mid := len(cards) / 2
	left := MergeSortDesc(cards[:mid])
	right := MergeSortDesc(cards[mid:])

	return merge(left, right)
}

func merge(left, right [][]int) [][]int{
	resSlice := make([][]int, 0, len(left) + len(right))
	
	for len(left) != 0 || len(right) != 0 {

		if len(left) == 0 {

			resSlice = append(resSlice, right[0])
			right = right[1:]

		}else if len(right) == 0 {

			resSlice = append(resSlice, left[0])
			left = left[1:]

		}else {

			valLeft := left[0][0] + left[0][1]
			valRight := right[0][0] + right[0][1]

			if valLeft > valRight{

				resSlice = append(resSlice, left[0])
				left = left[1:]
				
			}else {

				resSlice = append(resSlice, right[0])
				right = right[1:]

			}
		}

	}
	
	return resSlice

}

func playingDomino(cards [][]int, deck []int) interface{} {
	// Di sort, agar dapat membuang kartu domino dengan nominal yang tinggi terlebih dahulu
	cards = MergeSortDesc(cards)
	out := []int{}
	for _, card := range cards {
		if card[0] == deck[1] {
			out = append(out, card...)
			break
		}
	}
	if len(out) == 0 {
		return "tutup kartu"
	}
	return out
}

func main() {
	fmt.Println(playingDomino([][]int{{6,5},{3,4},{2,1},{3,3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6,5},{3,3},{3,4},{2,1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6,6},{2,4},{3,6}}, []int{5, 1}))
}