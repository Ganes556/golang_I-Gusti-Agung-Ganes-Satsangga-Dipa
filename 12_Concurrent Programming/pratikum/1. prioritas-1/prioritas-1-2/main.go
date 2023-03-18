package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		x := 3
		for i := 1; i <= 10; i++ {
			ch <- x * i
		}
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}

}