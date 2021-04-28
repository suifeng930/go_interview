package rpc_hello

import "net/rpc"

// 服务的名称 (为了避免名字冲突，我们在RPC服务的名字中增加了包路径前缀，这是RPC服务抽象的包路径，并非完全等价于Go语言的包路径)
const HelloServiceName = "path/to/pkg.HelloService"

// 服务要实现的详细方法列表
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 注册该类型服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
