package main

import (
	"sync"
	"time"
)

var rwMu sync.RWMutex
var count int

// 死锁
//func main() {
//
//	go StartHttpDebuger()
//	go RWA()
//	time.Sleep(2 * time.Second)
//	rwMu.Lock()
//	defer rwMu.Unlock()
//	count++
//	fmt.Println(count)
//}

func RWA() {
	rwMu.RLock()
	defer rwMu.Unlock()
	RWB()
}

func RWB() {
	time.Sleep(time.Second * 5)
	RWC()
}

func RWC() {
	rwMu.RLock()
	defer rwMu.Unlock()
}
