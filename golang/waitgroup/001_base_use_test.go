package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBaseUse(t *testing.T) {
	var wg sync.WaitGroup

	start := time.Now()
	wg.Add(5) // 表示有5个goroutine需要等待
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done() // goroutine 退出时调用Done表示该写出已结束
			time.Sleep(time.Second)
			fmt.Printf("%d done\n", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Since(start).Seconds())
}

// 0 done
// 3 done
// 1 done
// 2 done
// 4 done
// 1.000787733
