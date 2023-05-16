package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var count int

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				count++
			}
		}()
	}

	wg.Wait()

	// 没有加锁, 每次运行的结果都不一样
	fmt.Printf("count: %v\n", count)
}

// 使用 go run -race main.go 可以检测是否有并发问题
// 但是只有触发了 data race 之后, 才能检测到
