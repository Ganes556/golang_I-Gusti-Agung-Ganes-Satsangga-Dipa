package main

import "fmt"

func main() {
	
	var n int
	
	fmt.Print("Input: ")

	fmt.Scan(&n)

	for i:=0; i < n; i++{

		for j:=0; j < n-i; j++{	
			fmt.Print(" ")
		}
		
		for k:= 1; k < (i+1)*2; k++{
			fmt.Print("*")
		}
		
		fmt.Println()

	}

}