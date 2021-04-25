package main

import (
	"log"
	"unicode"
)

func main() {




	//str :="abcdedfg"
	//s, b := reverseString(str)
	//log.Println(s,b)

	log.Println("机器人坐标问题")

	log.Println(move("R2(LF)",0,0,Top))
}

const (
	Left =iota     //0
	Top
	Right
	Bottom
)

// 机器人坐标问题
func move(cmd string, x0 int, y0 int, z0 int) (x,y,z int) {
	x,y,z =x0,y0,z0

	repeat :=0
	repeatCmd :=""

	for _, s := range cmd {

		switch  {
		case unicode.IsNumber(s):
			repeat=repeat*10+(int(s)-'0')
		case s ==')':
			for i :=0;i<repeat;i++ {
				x,y,z=move(repeatCmd,x,y,z)
			}
			repeat=0
			repeatCmd=""
		case repeat>0 && s!='(' && s !=')':
			repeatCmd=repeatCmd+string(s)
		case s=='L':
			z=(z+1)%4
		case s=='R':
			z=(z-1+4)%4
		case s=='F':
			switch  {
			case z==Left ||z ==Right:
				x=x-z+1
			case z==Top||z==Bottom:
				y=y-z+2
			}
		case s=='B':
			switch  {
			case z==Left|| z==Right:
				x=x+z-1
			case z==Top||z==Bottom:
				y=y+z-2
			}
		}
	}
	return x,y,z
}
