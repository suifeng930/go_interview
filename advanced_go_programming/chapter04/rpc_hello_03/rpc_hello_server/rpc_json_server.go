package main

import (
	"go_interview/advanced_go_programming/chapter04/rpc_hello_01/rpc_01"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	err := rpc.RegisterName("HelloService", new(rpc_01.HelloService))
	if err!=nil {
		log.Fatal("Register server err: ",err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err!=nil {
		log.Fatal("Listen TCP error: ",err)
	}

	for  {
		conn, err := listener.Accept()
		if err!=nil {
			log.Fatal("Accept err: ",err)
		}
		// rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) 针对服务端的json编解码器。
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}

}
