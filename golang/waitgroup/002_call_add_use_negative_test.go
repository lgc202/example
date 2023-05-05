package waitgroup

import (
	"sync"
	"testing"
)

// Add 传入负数
func TestCallAddUseNegative(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(-1)
	wg.Add(-1)
}
