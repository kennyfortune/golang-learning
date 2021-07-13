package basiclearning

import "fmt"

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func example() {
	// 用于检查是否有实现某个接口
	var _ Person = (*Student)(nil)

	var p Person = &Student{
		name: "Tom",
		age:  18,
	}
	// 接口转为实例
	stu := p.(*Student)
	fmt.Println(stu.getName())
}
