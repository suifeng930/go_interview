package main

import (
	"fmt"
	"time"
)

func main() {

	cannel :=make(chan bool)
	go worker(cannel)
	time.Sleep(time.Second)
	cannel<-true
}

func worker(cannel chan bool)  {
	for  {
		select {
		default:
			fmt.Println("hello")
			//正常工作
			case <-cannel:
				// 退出
		}
	}
}

