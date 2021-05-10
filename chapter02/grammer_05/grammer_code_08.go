package main

import "fmt"

func main02() {

	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch) //panic: close of nil channel
	}()
	<-ch
	fmt.Println(count)
}
