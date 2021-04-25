## Go常见面试题目汇总(01)

[TOC]

### 1. 交替打印数字和字母

#### 问题描述

>使用两个`goroutine`交替打印序列，一个`goroutine`打印数字，另一个`goroutine`打印字母，最终效果如下：
>
>```sh
>12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
>```

#### 解题思路

问题很简单，使用`channel`来控制打印的进度。使用两个`channel`，来分别控制数字和字母的打印序列，数字打印完成后通过`channel`通知字母打印，字母打印完成后通知数字打印，然后周而复始的工作。

#### 代码实现

```go
func main() {

	letter ,number :=make(chan bool),make(chan bool)
	wait :=sync.WaitGroup{}

	go func() {
		i :=1
		for  {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter<-true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i:=0
		for  {
			select {
			case <-letter:
				if i>=strings.Count(str,"")-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i:i+1])
				i++
				if i>=strings.Count(str,"") {
					i=0
				}
				fmt.Print(str[i:i+1])
				i++
				number<-true
				break
			default:
				break
			}
		}
	}(&wait)
	number<-true
	wait.Wait()
}
```

#### 源码解析

这里用到了两个`channel`负责通知，`letter`负责通知打印字母的`goroutine`来打印字母，`number`用来通知打印数字的`goroutine`打印数字。`wait`用来等待字母打印完成后退出循环。

### 2.判断字符串中字符是否完全都不同

#### 问题描述

请实现一个算法，确定一个字符串的所有字符「是否全都不同」。这里我们要求「不允许使用额外的存储结构」。给定一个string,请返回一个bool值，`true`代表所有所有字符全都不同，`false`代表存在相同的字符。保证字符串中的字符为「ASCII」字符。字符串的长度小于等于「3000」

#### 解题思路

这里有几个重点，第一个是`ASCII`字符，`ASCII`字符一共有256个，其中128个是常用字符，可以在键盘上输入，128之后的是键盘上无法找到的。

然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的存储结构，且字符串小于等于3000.

如果允许其他额外存储结构，这个提供很好写，如果不允许的话，可以使用golang内置的方式实现。

#### 源码参考

* 通过`strings.Count`函数判断：

```go

// 通过 steings.count判断
// strings.Count  返回字符串s 中有杰哥不重复的sep 子串
// 如果subStr 是一个空字符串,则 count 返回(1+s)中的unicode代码点数
func isUniqueString	(s string) bool  {

	if strings.Count(s,"")>3000 {
		return false
	}

	for _, v := range s {
		if v>127 {
			return false
		}
		if strings.Count(s,string(v))>1 {
			return false
		}
	}
	return true
}
```



* 通过`strings.Index`和`strings.LastIndex`函数判断

```go

// strings.Index 返回s 中substr的第一个实例的索引，如果不存在返回-1
func isUniqueString2(s string) bool {

	if strings.Count(s,"")>3000 {
		return false
	}
	for k, v := range s {
		if v>127 {
			return false
		}
		if strings.Index(s,string(v)) !=k {
			return false
		}
	}
	return false
}
```

#### 源码解析

以上两种方法都可以实现这个算法。

* `strings.Count()`,可以用来判断在一个字符串中包含的substr的数量，如果substr大于1，说明存在重复字符串
* `strings.Index()和strings.LastIndex()`,用来判断指定字符串在另一个字符串的索引位置，分别是第一个发现位置和最后发现位置。

### 3.翻转字符串

#### 问题描述

> 请实现一个算法，在不使用「额外数据结构和存储空间」的情况下，翻转一个给定的字符串(可以使用单个过程变量)
>
> 给定一个string,请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

#### 解题思路

翻转字符串其实是将一个字符串以中间字符为轴，前后翻转，即将`str[len]`赋值给`str[0]`，将`str[0]`赋值给`str[len]`

#### 源码参考

```go
//  翻转字符串
func reverseString(s string) (string ,bool) {

	str :=[]rune(s)
	l :=len(str)
	if l>5000 {
		return s,false
	}
	for i :=0;i<l/2; i++ {
		str[i],str[l-1-i]=str[l-1-i],str[i]
	}
	return string(str),true
}
```

#### 源码解析

以字符串长度的1/2为轴，前后赋值

### 3.判断两个给定的字符串排序后是否一致

#### 问题描述

>给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。这里规定「大小写为不同字符」，且考虑字符串重点空格。给定一个`string s1`和`string s2`，请返回一个bool,代表两串是否重新排列后可相同。保证两串的长度都小于等于5000.

#### 解题思路

首先要保证字符串长度小于5000.之后只需要一次遍历s1中的字符在s2中是否都存在即可。

#### 源码参考

```go
// 判断两个字符串排序之后是否相等
func isRegroup(s1, s2 string) bool {
	
	sl1 :=len([]rune(s1))
	sl2 :=len([]rune(s2))

	if sl1 >5000 ||sl2 >5000 || sl1 !=sl2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1,string(v)) !=strings.Count(s2,string(v)) {
			return false
		}
	}
	return true
}
```

#### 源码解析

这里还是使用`strings.count()来判断字符是否一致

### 字符串替换问题

#### 问题描述

>请编写一个方法，将字符串中的空格全部替换为"%20".假定该字符串有足够的空间存放新增的字符，并且知道字符串的真实长度(小于等于1000)，同时保证字符串由「大小写的英文字母组成」。给定一个string为原始的串，返回替换后的string.

#### 解题思路

两个问题：第一个是只能是英文字母，第二个是替换空格。

#### 源码参考

```go
// 字符串替换问题
func replaceBlank(s string) (string,bool) {

	if len([]rune(s))>1000 {
		return s,false
	}
	for _, v := range s {
		if string(v)!=" "&& unicode.IsLetter(v)==false {
			return s,false
		}
	}
	return strings.Replace(s," ","%20",-1),true
}
```

#### 源码解析

这里使用了golang内置方法`unicode.IsLetter()`判断字符是否是字母，之后使用`strings.Replace()`来替换空格。

### 机器人坐标问题

#### 问题描述

>有一个机器人，给一串指令，L左转R右转，F前进一步，B后退一步，问最后机器人的坐标，最开始，机器人位于`(0,0)`,方向为正Y。可以输入重复指令n；比如`R2(LF)`这等于指令`RLFLF`.问最后机器人的坐标是多少？

#### 解题思路

这里的一个难点是解析重复指令。主要指令解析成功，计算坐标就简单了。

#### 源码参考

```go
func main() {
	log.Println("机器人坐标问题")
	log.Println(move("R2(LF)",0,0,Top))
}
const (
	Left =iota     //0
	Top
	Right
	Bottom
)
// 机器人坐标问题
func move(cmd string, x0 int, y0 int, z0 int) (x,y,z int) {
	x,y,z =x0,y0,z0
	repeat :=0
	repeatCmd :=""
	for _, s := range cmd {
		switch  {
		case unicode.IsNumber(s):
			repeat=repeat*10+(int(s)-'0')
		case s ==')':
			for i :=0;i<repeat;i++ {
				x,y,z=move(repeatCmd,x,y,z)
			}
			repeat=0
			repeatCmd=""
		case repeat>0 && s!='(' && s !=')':
			repeatCmd=repeatCmd+string(s)
		case s=='L':
			z=(z+1)%4
		case s=='R':
			z=(z-1+4)%4
		case s=='F':
			switch  {
			case z==Left ||z ==Right:
				x=x-z+1
			case z==Top||z==Bottom:
				y=y-z+2
			}
		case s=='B':
			switch  {
			case z==Left|| z==Right:
				x=x+z-1
			case z==Top||z==Bottom:
				y=y+z-2
			}
		}
	}
	return x,y,z
}

```

#### 	源码解析

这里使用三个值表示机器人当前的状况,分别是: x表示x坐标，y表示y坐标，z表示当前方向。L、R命令会改变值z, F、B命令会改变值x、y。值x、y的改变还受当前的z值影响。

如果是重复指令，那么将重复次数和重复的指令存起来递归调用即可。


