package rpc_hello

import "net/rpc"

//  HelloServiceClient 实现 HelloServiceInterface 接口
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// 新增客户端 HelloServiceClient 结构体，该类型必须满足HelloServiceInterface 接口，
// 这样客户端用户就可以直接通过接口对应的方法调用RPC 函数
type HelloServiceClient struct {
	*rpc.Client
}

// 实现 Hello 函数
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

// 提供 DialHelloService 函数，直接拨号 HelloService 服务
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}
