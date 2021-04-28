package main

import (
	"go_interview/advanced_go_programming/chapter04/rpc_hello_01/rpc_01"
	"log"
	"net"
	"net/rpc"
)

func main() {

	//    rpc.RegisterName() 调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在HelloService服务的空间之下。
	rpc.RegisterName("HelloService",new(rpc_01.HelloService))
	//   建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err!=nil {
		log.Fatal("ListenTCP error :",err)
	}
	conn, err := listener.Accept()
	if err!=nil {
		log.Fatal("Accept error:",err)
	}
	// 通过rpc.ServeConn(conn) 提供RPC服务
	rpc.ServeConn(conn)
}
