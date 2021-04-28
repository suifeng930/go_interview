package netrpcPlugin

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
)

type netrpcPlugin struct {
	// 内置 generator.Generator 成员
	*generator.Generator
}


// Name 返回插件的名字
func (p *netrpcPlugin) Name() string {
	return "netrpc"
}

func (p *netrpcPlugin)Init(g *generator.Generator)  {
	// 初始化的时候用参数g进行初始化，因此插件是从参数g 对象继承了全部的共有方法
	p.Generator=g

}

// GenerateImports 调用自定义的  genImportCode 生成导入代码
func (p *netrpcPlugin)GenerateImports(file *generator.FileDescriptor)  {
	if len(file.Service)>0 {
		p.genImportCode(file)
	}
}

// Generate  调用自定义的 genServiceCode 生成每个服务的代码
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor)  {

	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {

	p.P("// TODO: import code")
}

func (p *netrpcPlugin) genServiceCode(svc *descriptorpb.ServiceDescriptorProto) {

	p.P("// TODO: service code, Name="+svc.GetName())
}