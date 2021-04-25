package main

import (
	"log"
	"sync"
)

var Total struct{
	sync.Mutex
	value int
}

// 在worker的循环中，为了保证total.value +=i 的原子性，我们通过sync.Mutex 加锁和解锁来保证该语句在同一时刻只被一个线程访问。
func worker(wg  *sync.WaitGroup)  {

	defer wg.Done()
	for i :=0;i<=100;i++ {
		Total.Lock()
		Total.value+=i
		log.Println(Total.value)
		Total.Unlock()
	}
}
