# 前言
Go语言中的 WaitGroup 可以进行任务编排，比如我们有一个主任务在执行，执行到某一点时需要并行执行三个子任务，并且需要等到三个子任务都执行完后，再继续执行主任务。那我们就需要设置一个检查点，使主任务一直阻塞在这，等三个子任务执行完后再放行。
# 基本使用
我们先来个简单的例子，看下 WaitGroup 是怎么使用的。示例中使用 Add(5) 表示我们有 5个 子任务，然后起了 5个 协程去完成任务，主协程使用 Wait() 方法等待 子协程执行完毕，输出一共等待的时间。     
[点击查看源码](./001_base_use_test.go)
```go
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
```
# 原理分析
## 总览
WaitGroup 一共有三个方法：
```go
(wg *WaitGroup) Add(delta int)
(wg *WaitGroup) Done()
(wg *WaitGroup) Wait()
```
- Add 方法用于设置 WaitGroup 的计数值，可以理解为子任务的数量
- Done 方法用于将 WaitGroup 的计数值减一，可以理解为完成一个子任务
- Wait 方法用于阻塞调用者，直到 WaitGroup 的计数值为0，即所有子任务都完成  

正常来说，我们使用的时候，需要先确定子任务的数量，然后调用 Add() 方法传入相应的数量，在每个子任务的协程中，调用 Done()，需要等待的协程调用 Wait() 方法，状态流转如下图：

## 底层实现

# 易错点