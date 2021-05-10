package main

import (
	"sync"
)

func main04() {

	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(m.Len())

}
