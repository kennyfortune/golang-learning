package test_learning

import (
	"testing"
)

// TestMul 展示了普通注释的写法，使用表驱动执行实例
func TestMul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"pos", args{3, 4}, 12},
		{"neg", args{-3, 4}, -12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

type calcCase struct {
	a, b, want int
}

//TestAdd 展示了t.Hepler()的用法，注意需要传入调用方法的testing.T。当我们需要再测试中抽象公共方法，此方法需要调用t.Helper()，来确保它能被测试跟踪
func TestAdd(t *testing.T) {
	createAddTestCase(t, &calcCase{1, 2, 3})
	createAddTestCase(t, &calcCase{2, 4, 6})
	createAddTestCase(t, &calcCase{1, 2, 2})
}

func createAddTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := c.a + c.b; ans != c.want {
		t.Errorf("%d * %d expected %d, but %d got", c.a, c.b, c.want, ans)
	}
}
