package basiclearning

import (
	"context"
	"fmt"
	"time"
)

// select & chan

var stop chan bool

func reqTaskSC(name string) {
	for {
		select {
		case <-stop:
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func sc() {
	stop = make(chan bool)
	go reqTaskSC("worker1")
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(3 * time.Second)
}

// Context 并发控制
// 1. 通知子协程退出（正常退出，超时退出等）；
// 2. 传递必要的参数。

// context.WithCancel 创建可取消的 Context 对象，即可以主动通知子协程退出。

func reqTaskCWC(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func CWC() {
	// context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	// context.WithCancel(parent) 创建可取消的子 Context，同时返回函数 cancel。
	// 在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出。
	// 主协程中，调用 cancel() 函数通知子协程退出。
	ctx, cancel := context.WithCancel(context.Background())
	go reqTaskCWC(ctx, "worker1")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

// 如果需要往子协程中传递参数，可以使用 context.WithValue()。

type Options struct{ Interval time.Duration }

func reqTaskCWV(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func CWV() {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "options", &Options{1})
	go reqTaskCWV(vCtx, "worker1")
	go reqTaskCWV(vCtx, "worker2")
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

// 如果需要控制子协程的执行时间，可以使用 context.WithTimeout 创建具有超时通知机制的 Context 对象。
func CWTO() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	go reqTaskCWC(ctx, "worker1")
	go reqTaskCWC(ctx, "worker2")
	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}

// 超时退出可以控制子协程的最长执行时间，那 context.WithDeadline() 则可以控制子协程的最迟退出时间。
func reqTaskWDL(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", name, ctx.Err())
			return
		default:
			fmt.Println(name, "send request")
			time.Sleep(1 * time.Second)
		}
	}
}

func CWDL() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	go reqTaskWDL(ctx, "worker1")
	go reqTaskWDL(ctx, "worker2")

	time.Sleep(3 * time.Second)
	fmt.Println("before cancel")
	cancel()
	time.Sleep(3 * time.Second)
}
