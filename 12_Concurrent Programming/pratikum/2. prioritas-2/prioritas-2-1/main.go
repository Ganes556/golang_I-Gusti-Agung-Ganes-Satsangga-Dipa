package main

import (
	"fmt"
	"runtime"
	"sync"
)

func freqChar(text string, goroutines int) map[rune]int {
	freq := make(map[rune]int)

	var wg sync.WaitGroup

	sizeTextEachGoroutine := len(text) / goroutines
	
	for i := 0; i < goroutines; i++ { 

		wg.Add(1)

		start := i * sizeTextEachGoroutine
		end := start + sizeTextEachGoroutine
		
		go func() {
			for _, r := range text[start:end] {
				freq[r]++
			}
			defer wg.Done()
		}()
		wg.Wait()
	}

	return freq
}

func main() {
	runtime.GOMAXPROCS(2)
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
	for r, freq := range freqChar(text, 3){
		fmt.Printf("%c: %d\n", r, freq)
	}
}
