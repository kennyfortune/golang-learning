package test_learning

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	fmt.Printf("test begin\n")
}

func teardown() {
	fmt.Printf("test end\n")
}

func TestA(t *testing.T) {
	fmt.Printf("test a\n")
}

func TestB(t *testing.T) {
	fmt.Printf("test b\n")
}

//TestMain 这里展示了TestMain的用法，使用testing.M，go test此文件下测试用例将不会单独执行。同时这里也展示了setup、teardown的用法，可以用来封装测试资源的初始化和关闭。
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
