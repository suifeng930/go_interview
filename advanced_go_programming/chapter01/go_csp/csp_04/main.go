package main

import (
	"fmt"
	"go_interview/advanced_go_programming/chapter01/go_csp/csp_04/pubsub"
	"strings"
	"time"
)

func main() {

	p := pubsub.NewPublisher(100*time.Millisecond, 10)

	defer p.Close()
	all := p.Subscribe()

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello ,world ")
	p.Publish("hello ,golang ")

	go func() {
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	// 运行一定时间后退出

	time.Sleep(10 * time.Second)
}

// 在发布/订阅模型中，每条消息都会传递给多个订阅者。 发布者通常不会知道，也不关系哪一个订阅者正在接收主题消息。
// 订阅者和发布者可以在运行时动态添加，它们之间是一种松散的耦合关系，这使得系统的复杂性可以随时间的推移而增长。
//在现实生活中，像天气预报之类的应用就可以应用这种并发模式