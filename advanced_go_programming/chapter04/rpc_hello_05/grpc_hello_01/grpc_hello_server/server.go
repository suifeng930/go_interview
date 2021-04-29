package main

import (
	"context"
	pb "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_01/grpc_hello_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port=":1234"
)

type HelloServiceImpl struct {
	pb.UnimplementedHelloServiceServer   // 保持向前兼容
}

// 实现gRPC 的 接口
func (p *HelloServiceImpl)Hello(ctx context.Context,args *pb.String)(*pb.String,error)  {
	reply:=&pb.String{Value: "hello: -->"+args.GetValue()}
	return reply,nil
}

func main() {

	// 构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl服务
	pb.RegisterHelloServiceServer(grpcServer,new(HelloServiceImpl))

	listen, err := net.Listen("tcp", port)
	if err!=nil {
		log.Fatal(err)
	}
	// 在一个监听端口上提供gRPC服务
	err = grpcServer.Serve(listen)
	if err!=nil {
		log.Fatal("failed to server :",err)

	}

}
