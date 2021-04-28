## Go常见面试题汇总(02)

[TOC]

### 常见语法题目一

### 1.下面代码能运行吗？为什么。

```go

type Param map[string]interface{}
type Show struct {
	Param
}
func main1()  {
	s :=new(Show)             // new 无法初始化Show结构体中的param
	s.Param["RMB"]=10000     // panic: assignment to entry in nil map
}
```

**解析**

共发现两个问题：

1. `main()`函数不能加数字
2. `new()`关键字无法初始化`Show`结构体中的`Param`属性，所以直接对`s.Param`操作会报错`panic: assignment to entry in nil map`

------

### 2.请说出下面代码存在什么问题

```go
type student struct {
	Name string
}

func zhoujielun(v interface{})  {

	switch msg:=v.(type) {
	case *student,student:
		msg.Name  //Unresolved reference 'Name'
	}
}

```

**解析**

golang中有规定，`switch type`的`case T1`,类型列表只有一个，那么`v :=m.(type)`中的`v`的类型就是`T1`类型。

如果是`case T1,T2`，类型列表中有多个，那么`v`的类型还是多对应接口的类型，也就是`m`的类型。因此这里`msg`的类型还是`interface{}`,所以它没有`Name`这个字段，所以编译阶段就会报错。`Unresloved reference 'Name'`

------

### 3.写出打印的结果

```go
type People struct {
	name string `json:"name"`  // 小写属性不能被转码 包可见
}

func main() {

	js := `{ "name":"11l"}`

	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err!=nil {
		fmt.Println("err: ",err)
		return
	}
	fmt.Println("people: ",p)  // people:  {}
}
```

**解析**

按照golang的语法，小写开头的方法、属性或`struct`是私有的，同样，在`json`解码、转码的时候也无法实现私有属性的转换。

题目中是无法正常得到`people`的`name`属性的值，而且，私有属性`name`也不应该加`json`的标签。

------

### 4.下面的代码是有问题的，请说明原因

```go
type Person struct {
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("print: %v \n",p)
}

func main()  {

	p:=&Person{}
	p.String()  // fatal error: stack overflow
}
```

**解析**

在golang中`String() string`方法实际上是实现了`String`接口类型的，该接口定义在`fmt/print.go`文件中：

```
type Stringer interface {
		String() string
}
```

在使用`fmt`包中的打印方法时，如果类型实现了这个接口，会直接调用。而题目中打印`p`的时候会直接调用`p`实现的`String()`方法，然后产生了循环调用,报错`// fatal error: stack overflow`。

------

### 5.请找出下面代码的问题所在

```go
func main()  {

	ch :=make(chan int,1000)

	go func() {
		for i :=0;i<10;i++{
			ch<-i
		}
	}()

	go func() {
		a,ok:=<-ch
		if !ok {
			fmt.Println("close")
			return
		}
		fmt.Println("a: ",a)
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second*100)
}
```

**解析**

不满足golang的顺序一致性内存模型。在go中，同一个goroutine线程内部，顺序一致性的内存模型是得到保证的，但在不同的goroutine之间，并不满足顺序一致性的内存模型，需要通过明确定义的同步事件来作为同步的参考。

因此题中的goroutine之间的调度顺序是不确定的，有存在第一个`channel`的goroutine还未调用，或者已经调用但没有写完时，直接被main()调用了`close()`管道，可能导致写失败，进而出现`panic`错误。`panic: send on closed channel`	

------

### 6.请说明下面代码书写是否正确

```go
var value int32

func SetValue(delta int32)  {
	for  {
		v:=value
		if atomic.CompareAndSwapInt32(&value,v,(v+delta)) {
			break
		}
	}
}
```

**解析**

`atomic.CompareAndSwapInt32()`函数不需要循环调用。

------



### 7.下面的成勋运行后为什么会爆异常

```go
type Project struct {}

func (p *Project) deferError() {
	if err:=recover();err!=nil {
		fmt.Println("recover: ",err)
	}
}

func (p *Project)exec(msgChan chan interface{})  {
	for  msg := range msgChan {
		m :=msg.(int)
		fmt.Println("msg: ",m)
	}
}

func (p *Project) run(msgChan chan interface{})  {
	for  {
		defer p.deferError()
		go p.exec(msgChan)
		time.Sleep(time.Second*2)
	}
}
func (p *Project) Main()  {
	a :=make(chan interface{},100)
	go p.run(a)
	go func() {
		for  {
			a<- "1"
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second*100000000000000)
//	time.Sleep(time.Second*10000)
}

func main()  {
	p :=new(Project)
	p.Main()
}
```

**解析**

有以下问题：

* `time.Sleep()`的参数数值太大，超过了`1<<63-1`的限制
* 报错panic：`panic: interface conversion: interface {} is string, not int`
* `defer p.deferError()`需要在goroutine开始处调用，否则无法捕获`panic`

------

### 8.请说出下面代码哪里写错了

```go

func main()  {

	abc :=make(chan int, 1000)
	for i :=0;i<10;i++ {
		abc<-i
	}
	go func() {
		for  a := range abc {
			fmt.Println("a: ",a)
		}
	}()
	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second*10)
}
```

**解析**

协程可能还未启动，管道就关闭了。

------

### 9.请说出下面代码，执行时为什么会报错

```go
type Student struct {
	name string
}
func main()  {

	m :=map[string]Student{"people":{"zhoujielun"}}
	m["people"].name="wuyanzu"  // Cannot assign to m["people"].name
}
```

**解析**

map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。故如果需要修改map值，可以将`map`中的非指针类型`value`修改为指针类型，比如使用`map[string]*Student`

------

### 10.请说出下面的代码存在什么问题

```go
type query func(string) string

func exec(name string,vs ...query) string  {
	ch :=make(chan string)

	fn := func(i int) {
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func main()  {

	querys :=[]query{}
	querys1 := func(n string) string {
		return n+"func1"
	}
	querys2:= func(n string) string {
		return n+"func2"
	}
	querys3:= func(n string) string {
		return n+"func3"
	}
	querys4:= func(n string) string {
		return n+"func4"
	}
	querys=append(querys,querys1,querys2,querys3,querys4)
	ret :=exec("111", querys...)
	fmt.Println(ret)
}
```

**解析**

依据4个`goroutine`的启动后执行效率，很可能打印`111func4`,但其他的`111func（1，2，3）`也可能先执行，`exec()`只会返回一条消息。

------

### 11.下面这段代码为什么会卡死

```go
func main()  {

	var i byte
	go func() {
		for i  =0;i<=255;i++{ //Condition 'i<=255' is always 'true'
		}
	}()
	fmt.Println("Dropping mic")
	runtime.Gosched()
	runtime.GC()
	fmt.Println("Done")

}
```

**解析**

golang中，`byte`其实被`alias`到`uint8`上了。所以上main的`for`循环会始终成立，因为`i++ 到i=255`的时候会溢出，`i<=255`一定成立。也就是，`for`循环永远不会退出，因此上面的代码等价于：

```go
go func(){
	for{}   //死循环
}
```

正在被执行的`gorotine`发生以下情况时让出当前`goroutine`的执行权，并调度后面的`goroutine`执行：

* IO操作

* Channel阻塞

* System call

* 运行较长时间

  如果一个`goroutine`执行时间太长，`scheduler`会在其G对象上打上一个标记`preempt`,当这个`goroutine`内部发生函数调用的时候，会先主动检查这个标志，如果为`true`则会让出执行权。

  题中`main()`函数里启动的`goroutine`其实是一个没有IO阻塞、没有channel阻塞、没有system call、没有函数调用的死循环。

  也就是说，它无法主动让出自己的执行权，即使已经执行了很长时间，`scheduler`已经标志了`preempt`。而golang的GC动作是需要所有正在运行的`goroutine`都停止后进行的。因此程序会卡在`runtimee.GC()`等待所有协程退出。**这在go1.13及之前版本都会是这个结果，但go在1.14版本之后新增了异步抢占，为调度器提供了更强大的功能和控制。因此在go1.14及之后版本中，代码会正常执行完成。**

  **referrence**

  >抢占是Go调度器的重要组成部分，它分配了goroutine之间的运行时间。如果没有抢占，长期运行的goroutine将无止尽的消耗CPU以及阻止其他goroutine的调度运行。Go 1.14版本中新增了异步抢占，为调度器提供了更强大的功能和控制。
  >
  >https://zhuanlan.zhihu.com/p/138584406

