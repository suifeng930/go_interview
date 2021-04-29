package main

import (
	"context"
	pdd "go_interview/advanced_go_programming/chapter04/rpc_hello_05/grpc_hello_03/grpc_hello_publisher"
	"google.golang.org/grpc"
	"log"
)


func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(" conn err:", err)
	}
	defer conn.Close()

	client := pdd.NewPublishServiceClient(conn)

	_, err = client.Publish(context.Background(), &pdd.String{
		Value: "golang: hello GO",
	})
	if err != nil {
		log.Fatal(" client publish err:", err)
	}
	_, err = client.Publish(context.Background(), &pdd.String{
		Value: "docker: hello docker",
	})
	if err != nil {
		log.Fatal("client publish err:", err)
	}
}
