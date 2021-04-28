package main

import (
	"go_interview/advanced_go_programming/chapter04/rpc_hello_02/rpc_hello"
	"log"
	"net"
	"net/rpc"
)

func main() {

	err := rpc_hello.RegisterHelloService(new(rpc_hello.HelloService))
	if err != nil {
		log.Fatal("register HelloService err:", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Lister TCP error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		// 调用 rpc.ServeConn(conn) 函数提供RPC服务
		go rpc.ServeConn(conn)

	}

}
