package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 多次调用Done
func TestCallDoneMutiTimes(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("test")
		wg.Done()
		wg.Done()
	}()

	time.Sleep(time.Second)
	wg.Wait()
}
