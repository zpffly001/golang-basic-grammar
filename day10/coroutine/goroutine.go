package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var wg sync.WaitGroup

func runTask(id int) {
	// 推出一个减去1(表明该worker已经执行完毕，销毁，数量-1)
	defer wg.Done()

	fmt.Printf("task %d start..\n", id)
	// time.Sleep(2 * time.Second)
	fmt.Printf("task %d complete\n", id)
}

func asyncRun() {
	for i := 0; i < 10; i++ {
		go runTask(i + 1)
		// 启动worker，worker总数量+1
		wg.Add(1)
	}
}

func main() {
	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	asyncRun()
	// 阻塞，直到worker计数为0，即所有协程执行完毕
	wg.Wait()
}
