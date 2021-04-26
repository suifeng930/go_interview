package main

import (
	"log"
	"runtime"
)

func main() {

	var i byte

	go func() {
		for i = 0; i <= 254; i++ { //Condition 'i<=255' is always 'true'
		}
	}()
	log.Println("Dropping mic")

	runtime.Gosched()
	runtime.GC()
	log.Println("Done")

}
