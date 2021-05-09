package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type channelGO struct {
	sync.WaitGroup
	out chan int
}

func (c *channelGO) makeNumber() {

	go func() {
		defer c.WaitGroup.Done()
		for i := 0; i < 5; i++ {
			c.out <- rand.Intn(5)
		}
		close(c.out)
	}()
}

func (c *channelGO)printNumber()  {
	go func() {
		defer c.WaitGroup.Done()
		for i := range c.out {
			fmt.Println(i)
		}
	}()
	c.WaitGroup.Wait()
}