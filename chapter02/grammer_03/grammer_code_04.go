package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//
//	go func() {
//		// 1. 在这里需要写你的算法
//		// 2. 要求每秒钟调用一次proc函数
//		// 3. 要求程序不能退出
//		t := time.NewTicker(time.Second * 1) // 定义一个timeTicker 保证每秒调用
//		for {                                // 程序不能退出，死循环保证
//			select {
//			case <-t.C:
//				go func() {
//					defer func() {
//						if err := recover(); err != nil { // 使用recover() 捕获 panic
//							fmt.Println(err)
//						}
//					}()
//					proc()
//				}()
//			}
//		}
//	}()
//	select {}
//
//}
//
//func proc() {
//	panic("ok")
//}
