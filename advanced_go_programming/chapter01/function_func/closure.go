package main



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