package main

import "fmt"

func calc(index string, a, b int) int {

	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret

}

//10 1 2 3
//20 0 2 2
//2 0 2 2
//1 1 3 4