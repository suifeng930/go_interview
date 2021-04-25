package main


var Done=make(chan bool)
var msg string

func aGoroutine()  {
	msg="你好、世界"
	Done <-true
}

func bGoroutine() {

	msg =" close channel, 你好、世界"
	close(Done)
}