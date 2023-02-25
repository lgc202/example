# 1. golang 单元测试基础
(1) golang的测试文件必须以`*_test.go`结尾  
(2) golang测试函数签名为`func TestXXX(t *testing.T)`的形式  
(3) 例子
```go
// math.go

// Abs returns the absolute value of x.
func Abs(x float64) float64 {
    return math.Abs(x)
}

// Max returns the larger of x or y.
func Max(x, y float64) float64 {
    return math.Max(x, y)
}

// Min returns the smaller of x or y.
func Min(x, y float64) float64 {
    return math.Min(x, y)
}

// RandInt returns a non-negative pseudo-random int from the default Source.
func RandInt() int {
    return rand.Int()
}
```
编写的测试文件如下  
```go
// math_test.go

func TestAbs(t *testing.T) {
	got := Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

func TestMax(t *testing.T) {
	got := Max(1, 2)
	if got != 2 {
		t.Errorf("Max(1, 2) = %f; want 2", got)
	}
}
```
(4) go test常见参数
- -v，显示所有测试函数的运行细节
```shell
$ go test -v                                                                       
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
=== RUN   TestMax        
--- PASS: TestMax (0.00s)
PASS                     
ok      golang/utest/001base    0.307s
```
- -run < regexp>，指定要执行的测试函数
```shell
$ go test -v -run='TestA.*'
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok      golang/utest/001base    0.332s
```
- -count N，指定执行测试函数的次数
```shell
$ go test -v -run='TestA.*' -count=2
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
=== RUN   TestAbs
--- PASS: TestAbs (0.00s)
PASS
ok      golang/utest/001base    0.329s
```
# 2. 跳过某些测试
# 3. 表格驱动测试
# 4. 性能测试