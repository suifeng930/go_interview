package main

import (
	pbb "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_02/grpc_hello_stream"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	port = ":1234"
)

type HelloServiceImpl struct {
	pbb.UnimplementedHelloServiceServer // 保持向前兼容
}

//// 实现gRPC 的 接口
//func (p *HelloServiceImpl) Hello(ctx context.Context, args *pbb.String) (*pbb.String, error) {
//	reply := &pbb.String{Value: "hello: -->" + args.GetValue()}
//	return reply, nil
//}

//
func (p *HelloServiceImpl) Channel(stream pbb.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			log.Println(" server Channel recv err: ", err)
			return err
		}

		reply := &pbb.String{
			Value: "Hello-->Stream:" + args.GetValue(),
		}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {

	// 构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl服务
	pbb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	// 在一个监听端口上提供gRPC服务
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("failed to server :", err)

	}

}
