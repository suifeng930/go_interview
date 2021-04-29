package main

import (
	"context"
	pdd "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_03/grpc_hello_publisher"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal("dialing err:",err)
	}
	defer conn.Close()
	client := pdd.NewPublishServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &pdd.String{Value: "golang:"})
	if err != nil {
		log.Fatal("subscribe err:",err)
	}
	for  {
		reply, err := stream.Recv()
		if err!=nil {
			if err==io.EOF {
				break
			}
			log.Fatal("recv err:",err)
		}
		log.Println(reply.GetValue())
	}
}
