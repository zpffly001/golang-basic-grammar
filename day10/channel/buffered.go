package chnanel

import "fmt"

func senderV2(ch chan string, down chan struct{}, senderDown chan struct{}) {
	// ch <- v    // 发送值v到Channel ch中
	// v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v
	ch <- "hello"
	ch <- "this"
	ch <- "is"
	ch <- "alice"
	// 发送通话结束
	ch <- "EOF"

	// 同步模式等待recver 处理完成
	fmt.Println("sender wait ...")
	// 从down 接收端管道中获取数据
	<-down
	fmt.Println("sender down ...")

	// sender 退出
	senderDown <- struct{}{}

	// 处理完成后关闭channel
	close(ch)
}

// recver 循环读取chan里面的数据，直到channel关闭
func recverV2(ch chan string, down chan struct{}) {
	defer func() {
		down <- struct{}{}
	}()

	for v := range ch {
		// 处理通话结束
		if v == "EOF" {
			return
		}
		fmt.Println(v)
	}
}

func BufferedChan() {
	ch := make(chan string, 5)

	// 表明信道中存储的是结构体
	senderdown := make(chan struct{})
	recverdown := make(chan struct{})
	go senderV2(ch, recverdown, senderdown) // sender goroutine
	go recverV2(ch, recverdown)             // recver goroutine

	// 等待sender执行完成
	<-senderdown
}

// chan T          // 可以接收和发送类型为 T 的数据
// chan<- float64  // 只可以用来发送 float64 类型的数据
// <-chan int      // 只可以用来接收 int 类型的数据
