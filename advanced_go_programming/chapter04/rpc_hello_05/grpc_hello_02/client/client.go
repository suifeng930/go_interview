package main

import (
	"context"
	pbb "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_02/grpc_hello_stream"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pbb.NewHelloServiceClient(conn)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 客户端向服务端发送数据
	go func() {
		for {
			err := stream.Send(&pbb.String{Value: "hi"})
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	// 接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Println(reply.GetValue())
	}
}
