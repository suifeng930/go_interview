package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	// 调用 rpc.Dial() 拨号服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err!=nil {
		log.Fatal("dialing :",err)
	}

	var reply string
	// 调用 client.Call() 调用具体的RPC方法，调用Client.Call()时，第一个参数是用点号链接的RPC服务名字和方法名称，
	// 第二个和第三个参数分别是定义RPC（）方法的两个参数
	err = client.Call("HelloService.Hello", "hellosd", &reply)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
