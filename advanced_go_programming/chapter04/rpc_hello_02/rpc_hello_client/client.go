package main

import (
	"fmt"
	"go_interview/advanced_go_programming/chapter04/rpc_hello_02/rpc_hello"
	"log"
)

// 现在客户端用户不用再担心RPC方法名字或参数类型不匹配等低级错误的发生
func main() {

	client, err := rpc_hello.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("client dialing err: ", err)
	}
	var reply string
	// 唯一变化的是rpc_hello.HelloServiceName+".Hello"  替换了原来的 `HelloService.Hello()`
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println( reply)
}
