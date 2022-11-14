package cspmodel

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

// 发布订阅模型
func TestPubSubMode(t *testing.T) {
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	// 订阅所有
	all := p.Subscribe()

	// 通过过滤订阅一部分信息
	golang := p.SubscribeTopic(func(v interface{}) bool {
		// v是字符串类型
		if s, ok := v.(string); ok {
			// 订阅只包含"golang"的
			return strings.Contains(s, "golang")
		}
		return false
	})

	// 发布者 发布信息
	p.Publish("hello,   python!")
	p.Publish("godbybe, python!")
	p.Publish("hello,   golang!")

	// 订阅者查看消息(因为订阅了所有，因此会打印三次)
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	// 订阅者查看消息(因为只订阅了golang，因此只会打印一次)
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
