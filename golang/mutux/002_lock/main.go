package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex // Mutex 的零值是还没有 goroutine 等待的未加锁的状态

	var count int

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// 加锁后, 每次运行的结果都是 1000000
	fmt.Printf("count: %v\n", count)
}
