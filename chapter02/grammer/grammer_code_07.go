package main

import (
	"fmt"
	"time"
)

type Project struct {

}

func (p *Project) deferError() {
	if err:=recover();err!=nil {
		fmt.Println("recover: ",err)
	}
}

func (p *Project)exec(msgChan chan interface{})  {

	for  msg := range msgChan {
		m :=msg.(int)
		fmt.Println("msg: ",m)
	}
}

func (p *Project) run(msgChan chan interface{})  {

	for  {
		defer p.deferError()
		go p.exec(msgChan)
		time.Sleep(time.Second*2)
	}
}

func (p *Project) Main()  {

	a :=make(chan interface{},100)
	go p.run(a)

	go func() {

		for  {
			a<- "1"
			time.Sleep(time.Second)
		}
	}()
	//time.Sleep(time.Second*100000000000000)
	time.Sleep(time.Second*10000)

}

//func main()  {
//
//	p :=new(Project)
//	p.Main()
//}