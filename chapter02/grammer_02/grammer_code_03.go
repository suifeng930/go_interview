package main

import (
	"fmt"
	"sync"
)

// 闭包 + go runnext的问题
func sync_Wait_Group()  {

	wg :=sync.WaitGroup{}
	wg.Add(20)

	for i:=0;i<10;i++ {
		go func() {
			fmt.Println("i: ",i)    //
			wg.Done()
		}()
	}
	for j :=0;j<10;j++{
		go func(j int) {
			fmt.Println("j: ",j) //
			wg.Done()
		}(j)
	}
	wg.Wait()
}

//j:  9
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//i:  10
//j:  0
//j:  1
//j:  2
//j:  3
//j:  4
//j:  5
//j:  6
//j:  7
//j:  8