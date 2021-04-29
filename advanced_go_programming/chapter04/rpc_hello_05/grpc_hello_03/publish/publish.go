package main

import (
	"context"
	"github.com/docker/docker/pkg/pubsub"
	pdd "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_03/grpc_hello_publisher"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

const (
	port = ":1234"
)


type PubsubService struct {
	pub *pubsub.Publisher
	pdd.UnimplementedPublishServiceServer
}

func NewPubsubService() *PubsubService {
	return &PubsubService{pub: pubsub.NewPublisher(100*time.Millisecond, 10)}
}

func (p *PubsubService) Publish(ctx context.Context, arg *pdd.String) (*pdd.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pdd.String{}, nil
}

func (p *PubsubService) Subscribe(arg *pdd.String, stream pdd.PublishService_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})
	for v := range ch {
		if err := stream.Send(&pdd.String{Value: v.(string)}); err != nil {
			return err
		}
	}
	return nil
}


func main() {

	// 构造一个gRPC服务对象
	grpcServer := grpc.NewServer()
	// 然后通过gRPC插件生成的 RegisterHelloServiceServer 函数注册我们实现的 HelloServiceImpl服务
	pdd.RegisterPublishServiceServer(grpcServer, NewPubsubService())

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
