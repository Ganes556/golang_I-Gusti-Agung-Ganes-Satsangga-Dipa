package main

import "fmt"

func main() {
	var n int

	fmt.Print("Angka: ")
	fmt.Scan(&n)
	
	if n %2 == 0 {
		fmt.Println("Angka", n ,"adalah genap")
		return
	}

	fmt.Println("Angka", n ,"adalah ganjil")

}