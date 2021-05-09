package main

import "sync"

type threadSafeSet struct {
	s []interface{}
	sync.RWMutex
}

func (set *threadSafeSet)Iter() <-chan  interface{}{

	ch :=make(chan interface{})
	go func() {
		set.RLock()
		for elem := range set.s{
			ch<-elem
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}