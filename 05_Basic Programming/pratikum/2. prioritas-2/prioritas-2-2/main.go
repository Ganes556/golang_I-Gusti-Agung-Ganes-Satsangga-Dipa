package main

import "fmt"

func main() {
	
	var n int
	
	fmt.Print("Input: ")
	
	fmt.Scan(&n)
	
	fmt.Println("Output: ")

	for i:=1; i <= n; i++{

		for j:=1; j <= n-i; j++{	
			fmt.Print(" ")
		}
		for k:= 1; k <= i; k++{
			fmt.Print("* ")
		}
		fmt.Println()

	}

}