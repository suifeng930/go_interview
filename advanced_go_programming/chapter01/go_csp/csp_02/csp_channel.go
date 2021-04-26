package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	//noCacheChan()    //无缓存channel
	//CacheChan()      //带缓存channel

	//done :=make(chan int ,10)  // 带10个缓存
	//startChan(done)
	//reserveChan(done)

	var wg sync.WaitGroup

	// 开启N个后台打印线程
	for i :=0;i<10;i++{
		wg.Add(1)

		go func() {
			fmt.Println("你好、世界")
			wg.Done()
		}()
	}

	//等待N个后台线程完成
	wg.Wait()
}
// 正常运行： 你好、世界
// 对于从无缓存通道进行的接收，发生在对该通道进行的发送完成之前。
//因此，后台线程`<-done`接收操作完成之后，main()线程的`done<-1`发送操作才可能完成(从而退出main()、退出程序)，而此时打印工作已经完成了。
func  noCacheChan()  {


	done :=make(chan int)

	go func() {
		fmt.Println("你好、世界")
		<-done
	}()
	done<-1
}
// 对于带缓存的通道，对于通道的第K个接收完成操作发生在第K+C个发送操作完成之前，其中C是通道的缓存大小。
// 虽然通道是带缓存的，但是`main()`线程接收完成是在后台线程发送开始但还未完成的时刻，此时打印工作也是已经完成的
func CacheChan()  {
	done :=make(chan int,1) // 带缓存通道
	go func() {
		fmt.Println("你好，世界")
		done<-1
	}()
	<-done
}

func startChan( done chan int)  {
	// 开启n个后台打印线程
	for i:=0;i<cap(done);i++ {
		go func() {
			log.Println("你好、世界")
			done<-1
		}()
	}
}
func reserveChan(done chan int)  {
	// 等待N个后台线程完成
	for i:=0;i<cap(done);i++ {
		<-done
	}
}