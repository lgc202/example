package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

func doSomething2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("do  something")
}

// 正确使用, 在Wait之前调用Add
func TestAddBeforeWait(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go doSomething2(&wg)
	}

	wg.Wait()
	fmt.Println("main")
}
