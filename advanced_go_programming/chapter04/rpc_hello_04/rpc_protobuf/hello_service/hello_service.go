package hello_service

type HelloService struct {

}

// Hello 方法的输入参数和 输出参数均改用 protobuf 定义的String类型表示
func (p *HelloService) Hello(request *String ,reply *String)error  {
	reply.Value="Hello:-->"+request.Value
	return nil
}
