package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	//cannel := make(chan bool)
	//for i := 0; i < 10; i++ {
	//	go worker(cannel)
	//}
	//
	//time.Sleep(time.Second)
	//close(cannel)  // 使用close() 关闭一个通道来实现广播的效果

	cancel :=make(chan bool)
	var wg sync.WaitGroup
	for i :=0;i<10;i++{
		wg.Add(1)
		go Worker(&wg,cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}

func worker(cannel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello ")
		//正常工作
		case <-cannel:
			//退出
		}
	}
}

func Worker(wg *sync.WaitGroup,cannel chan bool)  {
	defer wg.Done()

	for  {
		select {
		default:
			fmt.Println("hello")
			//正常工作
			case <-cannel:
				//退出
				log.Println("cannel out")
				return
		}
	}
}