package main

import (
	"fmt"
	"sync"
)

func Kelipatan3SebanyakN(n int) int{
	var x int
	var wg sync.WaitGroup
	var mute sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(){
			// mutex.Lock akan mencegah variabel x agar hanya satu goroutine yang dapat mengaksesnya
			// dan mencegah terjadinya race condition saat variabel x diubah
			mute.Lock()
			x += 3
			mute.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return x
}
func main() {
	fmt.Println(Kelipatan3SebanyakN(1000))
}
