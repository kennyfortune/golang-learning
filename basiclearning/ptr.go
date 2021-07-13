package basiclearning

import "fmt"

/**
	指针

	Go三种指针:
	*类型: 普通指针类型，用于传递对象地址，不能进行指针运算。
	unsafe.Pointer: 通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）。
	uintptr: 用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。
			 uintptr is an integer type that is large enough to hold the bit pattern of any pointer.

	总结: unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。

**/

func pointerExample() {

	str := "Golang"
	var p *string = &str
	*p = "hello"
	fmt.Println(str)
}
