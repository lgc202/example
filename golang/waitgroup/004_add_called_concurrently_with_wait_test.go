package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

func doSomething(wg *sync.WaitGroup) {
	wg.Add(1)
	fmt.Println("do  something")
	defer wg.Done()
}

// Add 和 Wait 并发调用
func TestAddCalledConcurrentlyWithWait(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go doSomething(&wg)
	}

	wg.Wait()
	fmt.Println("main")
}
