package chnanel

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	Basic()
}

func TestBufferedChan(t *testing.T) {
	BufferedChan()
}

func TestSyncAB(t *testing.T) {
	SyncAB()
}

func TestDeadLockV1(t *testing.T) {
	ch := make(chan string)
	// send
	go func() {
		ch <- "hello"
	}()

	// receive	<-ch 从信道中接收数据
	{
		fmt.Println(<-ch)
	}
}

func TestDeadLockV2(t *testing.T) {
	ch := make(chan string, 1)
	// send		ch <- "hello"把数据放到信道中
	{
		ch <- "hello"
	}

	// receive		<-ch 从信道中接收数据
	{
		fmt.Println(<-ch)
	}
}

func TestRunTaskWithPool(t *testing.T) {
	RunTaskWithPool()
}
