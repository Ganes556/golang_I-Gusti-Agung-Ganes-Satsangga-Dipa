package main

import "fmt"

func caesar(offset int, input string) string {
	// your code here
	var cipher string
	for _, v := range input {
		if v == ' ' {
			cipher += " "
			continue
		}
		cipher += string((v - 'a' + int32(offset))%26 + 'a')
	}
	return cipher
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // cnvc
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
  fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}