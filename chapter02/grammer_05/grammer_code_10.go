package main

import "time"

var c = make(chan int)
var a int

func f() {
	a = 1
	time.Sleep(2*time.Second)
	<-c
}

func main05() {
	go f()
	c <- 0
	print(a)    // 1

}
