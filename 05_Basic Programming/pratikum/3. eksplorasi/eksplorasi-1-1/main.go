package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Apakah palindrome?")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Kata: ")

	scanner.Scan()

	var kata string = scanner.Text()
	
	fmt.Println("Captured:", kata)
	
	var revKata string
	for i := len(kata)-1; i >= 0; i--{
		revKata += string(kata[i])
	}
	
	if revKata == kata{
		fmt.Println("Palindrome")
		return
	}
	
	fmt.Println("Bukan Palindrome")
	
}