package main

import "fmt"

//  写已经关闭的chan
//func main() {
//
//	c := make(chan int, 3)
//	close(c)
//	c <- 1
//
//}

//panic: send on closed channel
//goroutine 1 [running]:
//main.main()

// 读已经关闭的chan

func main11() {

	fmt.Println("以下是数值的chan")
	ci := make(chan int, 3)
	ci <- 1
	close(ci)

	num, ok := <-ci

	fmt.Printf("读 chan 的协程结束： num= %v ,ok = %v \n", num, ok)

	num1, ok1 := <-ci
	fmt.Printf("再读 chan 的协程结束： num= %v ,ok = %v \n", num1, ok1)
	num2, ok2 := <-ci
	fmt.Printf("再再读 chan 的协程结束： num= %v ,ok = %v \n", num2, ok2)

	fmt.Println("以下是字符串chan")
	cs := make(chan string, 3)
	cs <- "aaaa"
	close(cs)
	str, ok := <-cs
	fmt.Printf("读 chan 的协程结束： str= %v ,ok = %v \n", str, ok)

	str1, ok1 := <-cs
	fmt.Printf("再读 chan 的协程结束： str= %v ,ok = %v \n", str1, ok1)

	str2, ok2 := <-cs
	fmt.Printf("再再读 chan 的协程结束： str= %v ,ok = %v \n", str2, ok2)

	type Mystruct struct {
		Name string
	}
	cstruct := make(chan Mystruct, 3)
	cstruct <- Mystruct{Name: "haha"}
	close(cstruct)

	strr, ok := <-cstruct
	fmt.Printf("读 chan 的协程结束： str= %v ,ok = %v \n", strr, ok)

	strr1, ok1 := <-cstruct
	fmt.Printf("再读 chan 的协程结束： str= %v ,ok = %v \n", strr1, ok1)

	strr2, ok2 := <-cstruct
	fmt.Printf("再再读 chan 的协程结束： str= %v ,ok = %v \n", strr2, ok2)

}
