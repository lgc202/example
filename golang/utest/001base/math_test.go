package test

import "testing"

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

// go test 常见参数
// -v，显示所有测试函数的运行细节
// -run < regexp>，指定要执行的测试函数
// -count N，指定执行测试函数的次数
