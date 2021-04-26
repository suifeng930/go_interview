package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("你好，世界")
		mu.Unlock()
	}()
	mu.Lock()

}

// fatal error: sync: unlock of unlocked mutex
