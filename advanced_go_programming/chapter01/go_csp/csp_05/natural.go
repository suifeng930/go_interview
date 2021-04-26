package main

import "fmt"

func main() {

	ch :=GenerateNatural()  // 自然数序列 2,3,4
	for i:=0;i<100;i++ {
		prime :=<-ch // 新出现的素数
		fmt.Printf("%v: %v\n",i+1,prime)
		ch =PrimeFilter(ch,prime)  // 基于新素数构造的过滤器
	}
}

//返回生成自然数序列的channel  2，3，4，5，...
func GenerateNatural() chan int {
	ch :=make(chan int)
	go func() {
		for i :=2;;i++{
			ch<-i
		}
	}()
	return ch
}

// 通道过滤器： 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int{
	out :=make(chan int)
	go func() {
		for  {
			if i :=<-in;i%prime!=0{
				out<-i
			}
		}
	}()
	return out
}