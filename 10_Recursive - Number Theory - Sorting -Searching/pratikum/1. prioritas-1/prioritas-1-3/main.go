package main

import (
	"fmt"
	"math"
)


func primeX(number int) int {
	if number == 1 {
		return 2
	}else {
		primeSeries := make([]int,0,number)
		// tambahkan 2, karena nilai prima pertama yang special
		primeSeries = append(primeSeries, 2)

		// karena bilangan prima yang genap cuma 2 sisanya ganjil
		i:= 3
		for len(primeSeries) != number {
			isPrime := true
			// cek tiap faktor ganjil dari i, hingga akar dari i. Apakah bisa membagi i
			for j:= 3; j <= int(math.Sqrt(float64(i))); j+=2{
				if i % j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeSeries = append(primeSeries, i)
			}
			i+=2
		}

		return primeSeries[number-1]
	}
}

// Driver code
func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
