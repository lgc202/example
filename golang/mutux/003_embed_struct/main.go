package main

import (
	"fmt"
	"sync"
)

// 常见用法是将 mutux 嵌入到结构体中

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("counter.Count: %v\n", counter.Count)
}
