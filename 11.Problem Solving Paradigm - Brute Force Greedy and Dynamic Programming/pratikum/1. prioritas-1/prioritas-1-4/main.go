package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {

	// karena x^2 + y^2 + z^2 = C, dan rentangan dari C 1-10000
	// sehingga bila masing-masing diakarkan, dengan C di asumsikan pada nilai maksimum = 10000,
	// maka akan mendapatkan bahwa hasil akar dari C = 100, sehingga nilai x,y,z tidak mungkin > 100
	max := 100
	
	for x:= 1; x < max; x++ {
		for y:= 1; y < max; y++ {
			// karena x + y + z == a, jadi z = a - x - y
			z := a - x - y
			distinct := x != y && x != z
			condition2 := x*y*z == b
			condition3 := x*x + y*y + z*z == c
			if z > 0 && distinct && condition2 && condition3{
				fmt.Println(x,y,z)
				return 
			}
		} 
	}
	fmt.Println("no solution")
	return
}

func main() {
	SimpleEquations(1,2,3)
	SimpleEquations(6,6,14)
}