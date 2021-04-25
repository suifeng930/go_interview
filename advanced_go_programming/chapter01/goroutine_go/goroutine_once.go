package main

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m sync.Mutex
	done uint32
}

func (o *Once)Do(f func())  {

	if atomic.LoadUint32(&o.done)==1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done==0 {
		defer atomic.StoreUint32(&o.done,1)
		f()
	}
}

type Singleton struct {}
var (
	Instances *Singleton
	once Once
)

func Instanced() *Singleton {
	once.Do(func() {
		Instances=&Singleton{}
	})
	return Instances
}
