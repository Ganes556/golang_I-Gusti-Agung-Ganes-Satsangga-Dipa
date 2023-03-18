package main

import (
	"fmt"
	"time"
)

func kelipatanX(x int) {
	for i:= 1;; i++{
		fmt.Println(x * i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	go kelipatanX(5)
	time.Sleep(30 * time.Second) // mencetak kelipatan 3 hingga 10 kali
}