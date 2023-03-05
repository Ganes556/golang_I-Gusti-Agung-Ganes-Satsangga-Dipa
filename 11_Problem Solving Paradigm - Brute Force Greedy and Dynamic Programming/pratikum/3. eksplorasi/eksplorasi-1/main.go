package main

import "fmt"

type romans struct {
	num int
	symbol string
}

func num2Rom(num int) string {
	if num <= 0 {
		return "Invalid Number"
	}

	romans := []romans{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, roman := range romans {
		for num >= roman.num {
			result += roman.symbol
			num -= roman.num
		}
	}

	return result
}

func main() {
	fmt.Println(num2Rom(4))    
	fmt.Println(num2Rom(9))    
	fmt.Println(num2Rom(23))   
	fmt.Println(num2Rom(2021)) 
	fmt.Println(num2Rom(1646)) 
}
