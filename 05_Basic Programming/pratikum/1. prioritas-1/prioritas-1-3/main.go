package main

import (
	"fmt"
)

func main() {
	
	var n int = -1
	fmt.Print("Masukan nilai: ")
	fmt.Scan(&n)
		
	if n < 0 || n > 100 {
		fmt.Println("Nilai Invalid")
	}else if n <= 34{
		fmt.Println("Nilai E")
	}else if n <= 49{
		fmt.Println("Nilai D")
	}else if n <= 64{
		fmt.Println("Nilai C")
	} else if n <= 79{
		fmt.Println("Nilai B")
	} else if n <= 100{
		fmt.Println("Nilai A")
	} 

}