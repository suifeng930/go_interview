package main

import (
	"context"
	pb "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_01/grpc_hello_service"
	"google.golang.org/grpc"
	"log"
)


// todo gRPC 和标准的RPC框架有一个区别，即gRPC生成的接口并不支持异步调用。
// todo 不过我么可以在多个goroutine之间安全的共享gRPC底层的HTTP/2链接，因此可以通过在另一个goroutine阻塞调用的方式异步调用
func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err!=nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 构造 NewHelloServiceClient 对象
	client := pb.NewHelloServiceClient(conn)

	reply, err := client.Hello(context.Background(), &pb.String{Value: "pb-->grpc-->hello"})
	if err!=nil {
		log.Fatal("could not greet: ",err)
	}
	log.Println(reply.GetValue())
}
