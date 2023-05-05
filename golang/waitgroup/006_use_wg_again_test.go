package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

// 没有等 Wait 返回，就重用 WaitGroup
func TestUseWaAgain(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("do something")
		wg.Done()
		wg.Add(1)
	}()

	wg.Wait()
}
