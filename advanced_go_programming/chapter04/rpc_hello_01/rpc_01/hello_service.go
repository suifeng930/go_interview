package rpc_01

type HelloService struct {

}

//  满足go语言的RPC规则：
//  1。 方法只能有两个可序列化的参数，其中第二个参数为指针类型
//  2。 并且返回一个error类型，同时必须是公开的方法
func (p *HelloService)Hello	(request string,reply *string)	error  {

	*reply="hello:"+request
	return nil
}
