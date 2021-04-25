package main

func main() {


	// 使用sync.MuTex  实现粗粒度的原子操作

	//var wg sync.WaitGroup
	//wg.Add(2)
	//go worker(&wg)
	//go worker(&wg)
	//wg.Wait()
	//fmt.Println(Total.value)

	//wg.Add(2)
	//go Worker(&wg)
	//go Worker(&wg)
	//wg.Wait()
	//fmt.Println(total)


	// 顺序一致性内存模型

	//go println("你好，世界")

	// 可以通过同步语句来个两个事件明确排序

	//done :=make(chan int)
	//
	//go func() {
	//	println("你好，世界")
	//	done <-1
	//}()
	//
	//<-done


	// 无缓冲的channel

	//go aGoroutine()
	//<-Done
	//println(msg)
	// 可保证打印`你好、世界`. 该程序首先对msg 进行写入，然后在Done 通道上发送同步信号，随后从Done接受到对应到同步信号，最后执行println()

	// close channel

	go bGoroutine()
	<-Done
	println(msg)

}

