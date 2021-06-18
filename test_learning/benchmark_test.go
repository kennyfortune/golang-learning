package test_learning

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

// BenchmarkHello 性能测试，必须以Benchmark开头
func BenchmarkHello(b *testing.B) {
	// 如果需要一些初始化的耗时操作，使用ResetTimer()重置定时器，防止前面的时间被计算进去
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println("hello")
	}
	fmt.Println("end")
}

// BenchmarkParallel 并行测性能
func BenchmarkParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
