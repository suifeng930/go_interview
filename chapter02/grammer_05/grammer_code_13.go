package main

import "fmt"

type Master struct {
	s string
}

// 这是上面提到的`在方法内部把局部变量指针返回` 的情况
func foo(s string) *Master {
	a :=new(Master)
	a.s=s
	return a  // 返回局部变量a,在C语言中是妥妥的野指针，但是在go则ok,但a会逃逸到堆上
}

func main13()  {
	a :=foo("hehehe")
	b:=a.s+" : hahaha"
	c:=b+" !"
	fmt.Println(c)
}

// 执行 go build -gcflags=-m main.go