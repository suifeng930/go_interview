package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"sync"
)

var mu sync.Mutex
var chain string

const (
	pprofAddr string =":7890"
)

func StartHttpDebuger()  {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/",http.HandlerFunc(pprof.Index))
	server  := &http.Server{Addr: pprofAddr, Handler: pprofHandler}
	go server.ListenAndServe()

}

// 死锁现象
//func main()  {
//
//	//
//	go StartHttpDebuger()
//	chain ="main"
//	A()
//	fmt.Println(chain)  // fatal error: all goroutines are asleep - deadlock!
//}

func A()  {
	mu.Lock()
	defer mu.Unlock()
	chain =chain+"---> A"
	fmt.Println(chain)
	B()
}
func B()  {
	chain=chain+" ---> B"
	fmt.Println(chain)
	C() //

}
func C()  {
	mu.Lock()   //
	defer mu.Unlock()
	chain=chain+" ---> C"
	fmt.Println(chain)
}