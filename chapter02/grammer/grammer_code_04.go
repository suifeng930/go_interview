package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("print: %v \n",p)

}
//
//func main()  {
//
//	p:=&Person{}
//	p.String()  // fatal error: stack overflow
//}