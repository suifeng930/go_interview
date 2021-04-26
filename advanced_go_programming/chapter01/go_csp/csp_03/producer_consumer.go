package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {

	ch :=make(chan int, 64) //成果队列

	go Producer(3,ch) // 生成3的倍数的序列
	go Producer(5,ch) // 生成5的倍数的序列
	go Consumer(ch)         // 消费生成的队列

	//运行一定事件后退出
	//time.Sleep(5*time.Second)
	// Ctrl+c 退出
	sig:=make(chan os.Signal,1)
	signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)
	fmt.Printf("quit (%v)\n",<-sig)
}

// 生产者： 生成factor整数倍的序列
func Producer( factor int,out chan<-int)  {
	for i :=0;;i++{
		out<-i*factor
	}
}

// 消费者
func Consumer(in <-chan int)  {
	for  v := range in {
		fmt.Println(v)
	}
}