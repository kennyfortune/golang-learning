package basiclearning

import (
	"fmt"
	"sync"
	"time"
)

// sync 用法

var wg sync.WaitGroup

// 例如我们希望并发下载 N 个资源，多个并发协程之间不需要通信
func downloadSync(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done()
}

func startSync() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go downloadSync("a.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")
}

var ch = make(chan string, 10)

func downloadChan(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url
}

func startChan() {
	for i := 0; i < 3; i++ {
		go downloadChan("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}
