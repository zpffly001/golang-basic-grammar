package chnanel

import (
	"fmt"
)

func A(send chan struct{}, recv chan struct{}) {
	defer wg.Done()

	msg := []string{"1", "2", "3"}

	count := 0
	for range recv {
		if count > 2 {
			return
		}

		fmt.Println(msg[count])
		count++
		// send <- struct {} {}：表示struct类型的值，该值也是空。它构造了一个struct {}类型的值，该值也是空。并把该空值放入信道。是不是有点刷新信道的意思？
		send <- struct{}{}
	}
}

func B(send chan struct{}, recv chan struct{}) {
	defer wg.Done()

	msg := []string{"A", "B", "C"}

	count := 0
	for range recv {
		if count > 2 {
			return
		}

		fmt.Println(msg[count])
		count++
		send <- struct{}{}
	}

}

func SyncAB() {
	a, b := make(chan struct{}), make(chan struct{})

	wg.Add(2)
	// 在A方法中a是发送方b是接收方。在B方法中b是发送方a是接收方；互为生产、消费方
	go A(a, b)
	go B(b, a)

	// 启动信号
	b <- struct{}{}

	wg.Wait()
}
