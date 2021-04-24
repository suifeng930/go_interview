package main

import "log"

func main() {

	inc := Inc()
	log.Println(inc)   // 43


	// 闭包的隐式问题
	ClosurePrint()
	// 闭包的隐式问题修复
	ClosurePrintV2()
	ClosurePrintV3()
}