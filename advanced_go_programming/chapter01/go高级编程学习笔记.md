## go_读书笔记

[toc]

### 数组、字符串和切片

> go语言数组、切片和字符串三者的关系

```sh
# 1. 在底层原始数据有着相同的内存结构；在上层，因为语法的限制而有着不同的行为表现。
# 2. 数组：一种值类型；
#					虽然数组的元素可以被修改，但数组本身的赋值和函数传参都是以整体复制的方式处理的。
# 3. 字符串：
#	         底层数据也是对应的字节数组，但字符串的只读属性禁止了在程序中对底层字节数组元素的修改。
#          字符串赋值只是复制了数据地址和对应的长度，而不会导致底层数据的复制。
# 4. slice（切片）
#          切片的结构和字符串结构类似；但解除了只读限制。
#          切片的底层数据虽然也是对应数据类型的数组，但每个切片还有独立的长度和容量信息，切片赋值和函数传参时也是将切片透信息
#          部分按值传递处理。因为切片头含有底层数据的指针，所以它的赋值也不会导致底层数据的复制。

```

其实go语言的**赋值**和**函数传参**规则很简单，除**闭包函数以引用的方式**对外部变量访问之外，其它赋值和函数传参**都以传值的方式处理**。

#### 数组

------

数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。**数组的长度**是数组类型的一部分，不同长度或不同类型组成的数组都是不同的类型。

```go
// 数组的定义方式
var a [3]int                 // 定义长度为3的int数组，元素全部为0
var b =[...]int{1,2,3}       // 定义长度为3int数组
var c =[...]int{2:3,1:2}     // 定义长度为3的int数组，{0,2,3}
var d =[...]int{1,2,4:5,6}   // 定义长度为6的int数组 {1,2,0,0,5,6}
```

go语言中数组是**值语义**。一个数组变量即表示整个数组。

可以将数组看作一个**特殊的结构体**，结构体的字段名对应数组的索引，同时结构体成员的数目是固定的(即len())。内置函数`len()`可用于计算数组的长度，`cap（）`可用于计算数组的容量。

```go
// 常见的遍历数组的方式

for i:=rang a {
  fmt.Printf("a[%d]  -->%d \n",i,a[i])
}

for i,v:=range b {
  fmt.Printf("b[%d]  -->%d \n",i,v)
}

for i:=0;i<len(c);i++ {
   fmt.Printf("c[%d]  -->%d \n",i,c[i])
}
```

用`for- range`方式遍历数组，性能可能会好一些，因为这种遍历保证不会出现数组越界的情况，每一轮迭代对数组元素的访问时可以省去对下标越界的判断。

#### 字符串

------

一个字符串是一个不可改变的字节序列，字符串通常用来表示人类可读的文本数据。

字符串的元素不可修改，是一个只读的字节数组。go语言的源代码中文本字符串通常被解释为采用了`UTF8`编码的`Unicode码点(rune)`序列。

------

go语言字符串的底层结构在`reflect.StringHeader`中定义：

```go
type StringHeader struct{
  Data uintptr   // 指向字符串的底层字节数组
  Len int        // 字符串的字节的长度
}
```

字符串其实是一个结构体，因此字符串的赋值操作也就是`refect.Stringheader`结构体的复制过程，并不会涉及底层字节数组的复制。

------

字符串虽然不是切片，但是支持切片操作，不同位置的切片底层访问的是同一块内存数据(因为字符串是**只读**的，所以相同的字符串面值常量通常对应同一个字符串常量)

```go
s :="hello,world"
hello :=s[:5]
world :=s[7:]
```

------

#### 切片

------

**切片(slice)**就是一种简化版的动态数组。

------

切片的结构定义，`reflect.SliceHeader`

```go
type SliceHeader struct {
  Data uintptr      // 切片指向的底层数组
  Len int           // 表示切片的长度， len<=cap
  Cap int           // 表示切片指向的内存空间的最大容量
}
```

切片的定义方式：

```go
var (

  a []int          // nil切片，  和nil相等，一般用来表示一个不存在的切片
  b =[]int{}       // 空切片，和nil不相等，一般用来表示一个空的集合
  c =[]int{1,2,3}  // 有三个元素的切片， len = cap =3
  d =c[:2]         // 有两个元素的切片， len=2  cap =3
  e =c[0:2:cap(c)] // 有两个元素的切片， len=2   cap =3
  f =c[:0]         // 有0个元素的切片，len=0  cap=3
  g =make([]int,3) // 有3个元素的切片，len  = cap = 3
  h =make([]int,2,3) // 有2个元素的切片，len  = 2 cap = 3
  i =make([]int,0,3) // 有0个元素的切片，len  = 0  cap = 3
)
```

**添加切片元素**

内置的泛型函数`append()`可以在切片的尾部追加n个元素

```go
var a []int
a =append(a,1)                    // 追加一个元素
a =append(a,1,2,3)                // 追加多个元素，手写解包方式
a =append(a,[]int{1,2,3}...)      // 追加一个切片，切片需要解包
```

不过要注意的是，在容量不足的情况下，`append()`操作会导致重新分配内存，可能导致巨大的内存分配和复制数据的代价。即使容量足够，依然需要用`append()`函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。

------

向切片的开头添加元素：

```go
var a =[]int{1,2,3}
a =append([]int{0},a...)                    // 在开头添加一个元素
a =append([]int{-3,-2,-1},a...)             // 在开头添加一个切片
```

在开头一般都会导致内存的重新分配，而且会导致已有的元素全部重新复制一次。因此在头追加元素的性能要比尾部追加性能差很多。

------

删除切片的元素

根据删除元素的位置，分为头删除、中间删除、尾部删除等情况。其中尾部删除性能最佳。

```go
a =[]int{1,2,3}
a =a[:len(a)-1]                         // 删除尾部1个元素
a =a[:len(a)-N]                         // 删除尾部n个元素
```

删除开头元素，可以采用直接移动数据指针的方式：

```go
a =[]int{1,2,3,4}
a =a[1:]                                // 删除开头1个元素
a =a[N:]                                // 删除开头N个元素
```

删除开头的元素也可以不移动数据指针，而将后面的数据向开头移动，可以用`append()`原地完成，(指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化)

```go
a =[]int{1,2,3,4}
a =append(a[:0],a[1:]...)                               // 删除开头1个元素
a =append(a[:0],a[N:]...)                               // 删除开头N个元素
```

也可以用`copy()`完成删除开头的元素：

```go
a =[]int{1,2,3,4}
a =a[:copy(a,a[1:])]                             // 删除开头1个元素
a =a[:copy(a,a[N:])]                               // 删除开头N个元素
```

对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用`append()`或`copy()`原地完成。

```go
a =[]int{1,2,3, ...}
a =append(a[:i], a[i+1]...)          // 删除中间1个元素
a =append(a[:i], a[i+N]...)          // 删除中间N个元素

a =a[:i+copy(a[:i], a[i+1]...)]          // 删除中间1个元素
a =a[:i+copy(a[:i], a[i+N]...)]          // 删除中间N个元素
```

------

切片高效操作的要点：要降低内存分配的次数，尽量保证`append()`操作不会超过`cap`的容量，降低触发内存分配的次数和每次分配内存的大小。

------

### 函数、方法和接口

------

函数对应操作序列，是程序的基本组成元素。go语言中的函数有具名和匿名之分：具名函数一般对应于包级别的函数，是匿名函数的一种特例。

当匿名函数**引用**了**外部作用域**中的**变量**时就成了**闭包函数**，闭包函数是函数式编程的核心。

**方法**是绑定到一个具体类型的特殊函数，Go语言中的**方法**是依托于类型的，必须在编译时**静态绑定**。

**接口**定义了**方法**的集合，这些方法依托于运行时的接口对象，因此接口对应的方法是在**运行时动态绑定**的。Go语言通过隐式接口机制实现了鸭子面向对象模型。

------

#### 函数

------

函数式第一类对象，可以将函数保存到变量中。函数主要有具名、匿名之分，包级别函数一般都是**具名函数**，**具名函数是匿名函数的一种特例**。**方法其实也是函数的一种**。

```go
// 具名函数
func Add(a,b int) int {
  return a+b
}
//匿名函数
var Add=func (a,b int) int{
  return a+b
}

```

------

Go语言中，函数可以有多个参数和多个返回值，参数和返回值**都是以传值的方式**和被调用者交换数据。此外，在语法上，还支持可变参数，可变数量的参数其实是一个切片类型的参数。

```go
// 多个参数和多个返回值
func swap(a,b int) (int,int){
  return b,a
}
// 可变数量的参数 
// more 对应的[]int 切片类型
func Sum(a int,more ...int) int{
  for _,v :=range more{
    a +=v
  }
  return a
}

```

------

当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同的结果：

```go
func main() {
  var a =[]interface{}{123,"abc"}
  Print(a...)  // 123. abc
  Print(a)     // [123,abc]
}
func Print(a ...interface{}){
  fmt.Println(a...)
}

```

------

不仅函数的参数可以有名字，也可以给函数的返回值命名。

如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值。

```go
func main() {

	inc := Inc()
	log.Println(inc)   // 43
}
//  如果返回值命名了，可以通过名字来修改返回值，也可以通过defer语句在return语句之后修改返回值。
//  其中defer语句延迟执行了一个匿名函数，因为这个匿名函数捕获了外部函数的局部变量v,这种函数我们一般称之为闭包。
//  闭包对捕获的外部变量并不是以传值方式访问，而是以引用方式访问。
func Inc() (v int) {
	defer func() {
		v++   // 在函数return 之后会执行++ 42+1= 43
	}()
	return 42
}

```

**闭包对捕获的外部变量并不是以传值方式访问，而是以引用方式访问。 ** **闭包**的这种以引用访问外部变量的行为可能会导致一些隐含的问题。

```go

func ClosurePrint()  {
	for i :=0; i<3; i++{
		defer func() {println(i)}()
	}

}
/**
* Output:
* 3
* 3
* 3
*/
// 解释： 因为是闭包，在for迭代语句中，每个defer语句延时执行的函数引用都是同一个i迭代变量，
//       在循环结束后这个变量的值为3,因此最终输出的结果都是3


```

 **解释**: 因为是**闭包**，在for迭代语句中，每个`defer`语句延时执行的函数引用**都是同一个i**迭代变量,在循环结束后这个变量的值为3,因此,最终输出的结果都是3。

```go
/**
* Output:
* 2
* 1
* 0
 */

//  修复思路：   在每轮迭代中为每一个defer语句 的闭包函数生成独有的变量。可以用下面两种方式：

func ClosurePrintV2()  {

	for i:=0;i<3;i++ {
		i:=i  // 定义一个循环体内局部变量i
		defer func() {println(i)}()
	}
}

func ClosurePrintV3()  {

	for i:=0;i<3;i++ {
		defer func(i int ) { // 通过函数传入： defer语句会马上对调用参数求值
			println(i)
		}(i)
	}
}
```

方法1是在循环体内部再定义一个局部变量，这样每次迭代defer语句的闭包函数捕获的都是不同的变量，这些变量的值对应迭代时的值。

方法2是将迭代变量通过闭包函数的参数传入，defer语句会马上对调用参数求值。

------

go语言中，如果以切片为参数调用函数，有时候会给人一种参数采用了传引用的方式的假象：因为在被调用函数内部可以修改传入的切片的元素。**其实，任何可以通过函数参数修改调用参数的情形，都是因为函数参数中显式或隐式传入了指针参数**。函数参数传值的规范更准确说是只针对数据结构中固定的部分传值，例如字符串或切片对应结构体中的指针和字符串长度结构体传值，但是并不包含指针间接指向的内容。将切片类型的参数替换为类似`reflect.SliceHeader`结构体就能很好理解**切片传值**的含义了：

```go
// 理解切片传值
func twice(x []int){
  for i:=range x{
    x[i]*=2
  }
}
type IntSliceHeader struct{
  Data []int
  Len int
  Cap int
}

func twice(x IntSliceHeader){
  for i:=0;i<x.Len;i++ {
    x.Data[i] *=2
  }
}
```

因为切片中的底层数组部分通过隐式指针传递(指针本身依然是传值的，但指针指向的却是同一份的数据)，所以被调用函数可以通过指针修改调用参数切片中的数据。

------

在go 1.4之前，**Go的动态栈**采用的是**分段式的动态栈**，简单说就是**采用一个链表来实现动态栈**，每一个链表的节点内存位置不会发生变化，但链表实现的动态栈对某些导致跨越链表不同节点的热点调用的性能影响较大，因为相邻的链表节点在内存位置一般不相邻，这**会增加CPU高速缓存面中失败的概率。**

为了解决热点调用的CPU缓存命中问题，Go1.4之后改用连续的动态栈实现，简单说是**采用一个类似动态数组的结构来表示栈**。

------

#### 方法

------

**方法**一般是面向对象编程的一个特性，Go语言的方法是关联到类型的，这样可以在编译阶段完成方法的静态绑定。一个面向对象的程序会用方法来表达其属性对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

面向对象编程更多的只是一种思想，Go语言中方式示例：

```go
//关闭文件
func (f *File) Close() error {
  // ...
}
// 读文件操作
func (f *File) Read(int64 offset, data []byte) int {
  // ...
}
```

Go语言中，对于给定的类型，每个方法的名字必须是唯一的，同时方法和函数一样也不支持重载。

------

Go语言不支持传统面向对象中的继承特性，而是以自己特有的组合方式支持了方法的继承。Go语言中，通过在结构体内置匿名的成员来实现继承：

```go
type Cache struct {
  m map[string]string
  sync.Mutex
}

func (p *Ceche) Lookup(key string) string{
  p.Lock()
  defer p.Unlock()
  return p.m[key]
}
```

Ceche结构体类型通过嵌入一个匿名的`sync.Mutex`来继承它的方法`Lock()`和`UnLock()`。但是在调用`p.Lock()`和`p.Unlock()`时，p并不是`Lock()`和`UnLock()`方法的真正接收者，而是会将它们展开为`p.Mutex.Lock()`和`p.Mutex.UnLock()`调用。这种展开是在编译期间完成的，并没有运行时代价。

------

#### 接口

------

Go的接口类型是对其它类型行为的抽象 和概括，因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让对象更加灵活和更具有适应能力。Go语言的接口类型是延时绑定，可以实现类似虚函数的多态功能。

------

### 面向并发的内存模型

------

常见的并行编程有多种模型，主要有多线程、消费传递等。从理论上来看，**多线程**和基于**消息传递**等并发编程是等价的。Go语言是基于消息并发模型的集大成者，它将基于CSP模型的并发编程内置到了语言中，通过一个go关键字就可以轻易地启动一个`goroutine`,Go语言的`goroutine`之间是共享内存的。

------

#### Goroutine和系统线程

------

`goroutine`是go语言特有的并发体，是一种轻量级的线程，由go关键字启动。在真实的Go语言的实现中，`goroutine`和系统线程也不是等价的。尽管两者的区别实际上只是一个量的区别，但正是这个量变引发了Go语言并发编程的飞跃。

------

首先，每个系统级线程都会有一个固定大小的栈(默认2MB),这个栈主要是用来保存函数递归调用时的参数和局部变量。固定了栈的大小导致了两个问题：

* 对于很多只需要很小的栈空间的线程是一个巨大的浪费；
* 对于少数需要巨大栈空间的线程又面临栈溢出的风险。

针对者两个问题的解决方案是：

* 降低固定的栈大小，提升空间的利用率；
* 增大栈的大小以允许更深的函数递归调用；

但两者是无法兼得的。相反，goroutine解决了这一囧境。

------

一个`goroutine`会以一个很小的栈启动(可能是2kb/4kb)，当遇到深度递归导致当前栈空间不足时，`goroutine`会根据需要动态的伸缩栈的大小(主流实现中栈的最大空间可达1GB)。因为启动的代价很小，所以我们可以轻易的启动成千上万个`goroutine`.

------

go调度器的工作原理和内核的调度是相似的，但这个调度器之关注单独的go程序中的goroutine.

`goroutine`采用的是半抢占式的协作调度，只有在当前`goroutine`发生阻塞时才会导致调度；同时发生在用户态，调度器会根据具体函数只保存必要的寄存器，切换的代价要比系统线程低很多。

------

运行时有一个`routime.GOMAXPROCS`变量，用于控制当前运行正常非阻塞`Goroutine`的系统线程数目。

------

#### 原子操作

------


