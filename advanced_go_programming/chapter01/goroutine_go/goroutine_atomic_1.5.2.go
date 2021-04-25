package main

import (
	"log"
	"sync"
	"sync/atomic"
)

var total uint64

func Worker(wg *sync.WaitGroup)  {

	defer wg.Done()
	var i uint64
	for i =0;i<=100;i++{
		atomic.AddUint64(&total,i)
		log.Println(total)
	}
}
