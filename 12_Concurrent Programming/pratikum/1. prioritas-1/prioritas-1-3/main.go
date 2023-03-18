package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		for i:= 1; i <= cap(ch); i++ {
			ch <- 3 * i
		}
		defer close(ch)
	}()
	

	for c := range ch {
		fmt.Println(c)
	}

}