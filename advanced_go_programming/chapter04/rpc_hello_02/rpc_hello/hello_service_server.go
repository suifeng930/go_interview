package rpc_hello

// 定义 HelloService 结构体，并实现 HelloServiceInterface 接口
type HelloService struct{}

// 实现 Hello 函数
func (p *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:->" + request
	return nil
}
