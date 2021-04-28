package main

import "fmt"

func defer_call()  {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	panic("触发异常")

}

//打印后
//打印中
//打印前
//panic: 触发异常
//goroutine 1 [running]:
//main.defer_call()