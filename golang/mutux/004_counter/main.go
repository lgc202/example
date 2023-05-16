package main

import (
	"fmt"
	"sync"
)

// 把获取锁、释放锁、计数加一的逻辑封装成一个方法，对外不需要暴露锁等逻辑

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("counter.Count: %v\n", counter.Count())
}
