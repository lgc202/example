package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

func TestCopyWg(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		// 这里没使用指针，wg状态一直不会改变，导致 Wait 一直阻塞
		go func(wg sync.WaitGroup) {
			fmt.Println("do  something")
			defer wg.Done()
		}(wg)
	}

	wg.Wait()
	fmt.Println("main")
}
