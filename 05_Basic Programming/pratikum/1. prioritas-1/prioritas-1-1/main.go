package main

import "fmt"

func main() {
	var T, a, b float64
	fmt.Print("Tinggi trapesium: ")
	fmt.Scan(&T)
	fmt.Print("Sisi sejajar `a`: ")
	fmt.Scan(&a)
	fmt.Print("Sisi sejajar `b`: ")
	fmt.Scan(&b)

	var L = 0.5 * T * (a+b)
	
	fmt.Println("Luas trapesium:", L)
	
}