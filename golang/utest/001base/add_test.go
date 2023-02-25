package _01base

import "testing"

func TestAdd(t *testing.T) {
	want := 3
	got := add(1, 2)
	if got != want {
		t.Errorf("expect: %d, got: %d", want, got)
	}
}

// go test 的选项
// -v，显示所有测试函数的运行细节
// -run < regexp>，指定要执行的测试函数
// -count N，指定执行测试函数的次数
