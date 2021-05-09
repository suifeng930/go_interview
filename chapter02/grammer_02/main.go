package main

import "fmt"

func main() {

	// 1。 写出下面代码输出内容
	//defer_call()

	// 2. 以下代码有什么问题，说明原因
	//paseStudent := pase_student()
	//
	//log.Println(paseStudent)

	// 3. 下面的代码会输出什么，并说明原因
	//runtime.GOMAXPROCS(1)
	//sync_Wait_Group()

	// 4. 下面代码会输出什么
	//t :=Teacher{}
	//t.ShowA()
	//t.ShowB()

	// 5.下面代码会触发异常吗？ 请详细说明
	//
	//runtime.GOMAXPROCS(1)
	//channel_select()
	// 6.下面代码输出什么

	//a :=1
	//b :=2
	//defer calc("1",a,calc("10",a,b))
	//a =0
	//defer calc("2",a,calc("20",a,b))
	//b=1

	// 7. 请写出以下输入内容
	//Print_code()

	// 10. 以下代码能编译过去吗？ 为什么
	//var Per Person=Student{}
	//think:="bitch"
	//fmt.Println(Per.Speak(think))

	//11 . 以下代码打印什么内容，为什么
	if live()==nil {
		fmt.Println("AAAAA")
	}else {
		fmt.Println("BBBBBB")
	}
}
