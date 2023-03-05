package main

import "fmt" 

func MergeSort(productPrice []int) []int {
	if len(productPrice) < 2 {
		return productPrice
	}

	mid := len(productPrice) / 2
	left := MergeSort(productPrice[:mid])
	right := MergeSort(productPrice[mid:])
	
	mergedArr := make([]int, 0, len(left)+len(right))
	
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0{
			mergedArr = append(mergedArr, right...)
			break
		}

		if len(right) == 0{
			mergedArr = append(mergedArr, left...)
			break
		}

		if left[0] < right[0]{
			mergedArr = append(mergedArr, left[0])
			left = left[1:]
		}else {
			mergedArr = append(mergedArr, right[0])
			right = right[1:]
			
		}
	}
	return mergedArr

}


func MaximumBuyProduct(money int, productPrice []int) {
	// your code here
	productPrice = MergeSort(productPrice)
	
	canBuy := 0
	for _, price := range productPrice {
		if money >= price {
			money -= price
			canBuy++
			continue
		}
		break
	}
	fmt.Println(canBuy)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})      // 3
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}) // 4
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})   // 4
  MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})           // 1
  MaximumBuyProduct(0, []int{10000, 30000})                        // 0
}