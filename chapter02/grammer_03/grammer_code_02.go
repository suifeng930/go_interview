package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {

	m.rmx.Lock() //nil
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {

		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {

	m.rmx.Lock()
	defer m.rmx.Unlock()

	item, ok := m.c[key]
	if !ok {
		time.Sleep(timeout)
	}
	if item == nil {
		fmt.Println("cannot  find value")
		return item
	} else {
		if item.isExist {
			if item.ch != nil {
				close(item.ch)
				item.ch = nil
			}
		}
		fmt.Println(item.value)
		return item.value
	}
}
